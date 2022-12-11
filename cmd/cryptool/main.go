package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/manifoldco/promptui"
	"github.com/nextdotid/creator_suite/config"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util/dare"
	"github.com/nextdotid/creator_suite/util/decrypt"
	"github.com/nextdotid/creator_suite/util/encrypt"
	"github.com/nextdotid/creator_suite/util/ipfs"
	"github.com/urfave/cli"
)

const (
	codeOK     int = iota // exit successfully
	codeError             // exit because of error
	codeCancel            // exit because of interrupt
)

var (
	cleanChan chan<- int // initialized in init
	cleanFn   = make([]func(int), 0, 3)

	ENCRYPT = "encrypt"
	DECRYPT = "decrypt"
	AES     = "AES"
	ECC     = "ECC"

	RawPassword       = "Raw Password"
	EncryptedPassword = "Encrypted Password"

	YES   = "Yes"
	NO    = "No"
	TRUE  = "True"
	FALSE = "False"
)

func main() {
	cleanChan := make(chan int, 1)
	go func() {
		code := <-cleanChan
		for _, f := range cleanFn {
			f(code)
		}
		os.Exit(code)
	}()
	config.Init()
	model.Init()
	err := InitCmdTool()
	if err != nil {
		fmt.Printf("err = %v\n", err)
		panic(err)
	}
}

func InitCmdTool() error {
	app := cli.NewApp()
	app.Name = "cryptool"
	app.Usage = "Crypto Tools for CreatorSuite"
	app.Version = "1.0.0.rc1"
	app.Action = func(ctx *cli.Context) error {
		fmt.Printf("\033[1;36;40m%s-%s: %s\033[0m\n", app.Name, app.Version, app.Usage)
		return nil
	}
	app.Commands = CryptoolCommands
	return app.Run(os.Args)
}

var CryptoolCommands = []cli.Command{
	{
		Name:    ENCRYPT,
		Aliases: []string{"encrypt", "en"},
		Usage:   "encrypt a file",
		Flags:   nil,
		Action: func(ctx *cli.Context) error {
			return commandAction(ctx, ENCRYPT)
		},
	},
	{
		Name:    DECRYPT,
		Aliases: []string{"decrypt", "de"},
		Usage:   "decrypt a file",
		Flags:   nil,
		Action: func(ctx *cli.Context) error {
			return commandAction(ctx, DECRYPT)
		},
	},
	{
		Name:    ECC,
		Aliases: []string{"ecc"},
		Usage:   "ecies for encrypt/decrypt password",
		Flags:   nil,
		Action: func(ctx *cli.Context) error {
			return commandAction(ctx, ECC)
		},
	},
}

func EccEncryptAction(ctx *cli.Context, content string, pubkey []byte) (string, error) {
	encryptDataByte, err := encrypt.EciesEncrypt([]byte(content), pubkey)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(encryptDataByte), nil
}

func EccDecryptAction(ctx *cli.Context, content string, privkey []byte) (string, error) {
	encryptDataByte, err := hexutil.Decode(content)
	if err != nil {
		return "", err
	}
	dataByte, err := decrypt.EciesDecrypt(encryptDataByte, privkey)
	if err != nil {
		return "", err
	}
	return string(dataByte), nil
}

func AesEncryptAction(ctx *cli.Context, aesKey string, src, dst *os.File) {
	key, err := encrypt.DeriveKey([]byte(aesKey), src, dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m %v\033[0m\n", err)
		os.Exit(codeError)
	}
	cfg := dare.Config{Key: key}
	if _, err := encrypt.AesEncrypt(src, dst, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m %v\033[0m\n", err)
		os.Exit(codeError)
	}
	cleanFn = append(cleanFn, func(code int) {
		dst.Close()
		if code != codeOK { // remove file on error
			os.Remove(dst.Name())
		}
	})
}

func AesDecryptAction(ctx *cli.Context, aesKey string, src, dst *os.File) {
	key, err := decrypt.DeriveKey([]byte(aesKey), src, dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m %v\033[0m\n", err)
		os.Exit(codeError)
	}
	cfg := dare.Config{Key: key}
	if _, err := decrypt.AesDecrypt(src, dst, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m %v\033[0m\n", err)
		os.Exit(codeError)
	}
	cleanFn = append(cleanFn, func(code int) {
		dst.Close()
		if code != codeOK { // remove file on error
			os.Remove(dst.Name())
		}
	})
}

func commandAction(ctx *cli.Context, action string) error {
	var err error
	switch action {

	case ENCRYPT:
		err := encryptPipeline(ctx)
		if err != nil {
			return err
		}
	case DECRYPT:
		err := decryptPipeline(ctx)
		if err != nil {
			return err
		}
	case ECC:
		err := eccPipeline(ctx)
		if err != nil {
			return err
		}
	default:
		err = fmt.Errorf("unsupported action [%s]", action)
	}
	return err
}

func encryptPipeline(ctx *cli.Context) error {
	fmt.Printf("Encrypting and publishing files to ipfs\n")
	prompt := promptui.Prompt{
		Label: "\033[1;36;40m Origin File (Input)\033[0m",
	}
	input, err := prompt.Run()
	if err != nil {
		return err
	}

	prompt = promptui.Prompt{
		Label: "\033[1;36;40m Encrypt File (Output)\033[0m",
	}
	output, err := prompt.Run()
	if err != nil {
		return err
	}

	in, out, err := parseIO(input, output)
	if err != nil {
		return err
	}
	pswd, err := getAesKey()
	if err != nil {
		return err
	}

	// encrypt content
	AesEncryptAction(ctx, pswd, in, out)
	fmt.Printf("\033[1;32;40m Encrypt content finished! %s\033[0m\n", out.Name())

	keyID, err := saveAesPswd(pswd)
	if err != nil {
		return err
	} else {
		fmt.Printf("\033[1;37;44m Password saved. [key id is %d ]\033[0m\n", keyID)
	}

	flag, err := parseBooleanSelection("Do you want to publish your file to IPFS?")
	if err != nil {
		return err
	}
	if !flag {
		fmt.Printf("\033[1;32;40mFinished!\033[0m\n")
		os.Exit(codeOK)
	}

	cfg, err := getIpfsConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m%v\033[0m\n", err)
		return err
	}

	// upload content
	IpfsUpload(cfg, out.Name())
	return nil
}

func decryptPipeline(ctx *cli.Context) error {
	fmt.Printf("Download and Decrypting files from ipfs\n")
	flag, err := parseBooleanSelection("Do you want to download your file from IPFS?")
	if err != nil {
		return err
	}
	input := ""
	parsePswd := ""
	if flag {
		cfg, err := getIpfsConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[1;31;40m%v\033[0m\n", err)
			return err
		}
		prompt := promptui.Prompt{
			Label: "\033[1;36;40m IPFS Location Url (/ipfs/${cid})\033[0m",
		}
		locationUrl, err := prompt.Run()
		if err != nil {
			return err
		}

		prompt = promptui.Prompt{
			Label: "\033[1;36;40m Download File\033[0m",
		}
		input, err = prompt.Run()
		if err != nil {
			return err
		}
		IpfsDownload(cfg, locationUrl, input)
	} else {
		prompt := promptui.Prompt{
			Label: "\033[1;36;40m Origin File (Input)\033[0m",
		}
		input, err = prompt.Run()
		if err != nil {
			return err
		}
	}

	prompt := promptui.Prompt{
		Label: "\033[1;36;40m Decrypt File (Output)\033[0m",
	}
	output, err := prompt.Run()
	if err != nil {
		return err
	}

	in, out, err := parseIO(input, output)
	if err != nil {
		return err
	}

	promptSelect := promptui.Select{
		Label: "\033[1;36;40m Choose password mode\033[0m",
		Items: []string{RawPassword, EncryptedPassword},
	}
	_, mode, err := promptSelect.Run()
	if err != nil {
		return err
	}
	if mode == RawPassword {
		parsePswd, err = getAesKey()
		if err != nil {
			return err
		}
	} else {
		encryptedPswd, err := getEncryptedDecryptionKey()
		if err != nil {
			return err
		}
		privKey, err := getPrivateKey()
		if err != nil {
			return err
		}
		parsePswd, err = EccDecryptAction(ctx, encryptedPswd, privKey)
		if err != nil {
			return err
		} else {
			fmt.Printf("\033[1;32;40m Decrypt encrypted password finished!\033[0m\n")
		}
	}

	// decrypt content
	AesDecryptAction(ctx, parsePswd, in, out)
	fmt.Printf("\033[1;32;40m Encrypt content finished! %s\033[0m\n", out.Name())
	return nil
}

func eccPipeline(ctx *cli.Context) error {
	promptSelect := promptui.Select{
		Label: "\033[1;36;40m Choose action\033[0m",
		Items: []string{ENCRYPT, DECRYPT},
	}
	_, mode, err := promptSelect.Run()
	if err != nil {
		return err
	}
	if mode == ENCRYPT {
		pswd, err := getAesKey()
		if err != nil {
			return err
		}
		pubKey, err := getPublicKey()
		if err != nil {
			return err
		}
		encryptedKey, err := EccEncryptAction(ctx, pswd, pubKey)
		if err != nil {
			return err
		}
		fmt.Printf("\033[1;32;40m Encrypted Password: %s\033[0m\n", encryptedKey)
		os.Exit(codeOK)
	} else {
		content, err := getEncryptedDecryptionKey()
		if err != nil {
			return err
		}
		privKey, err := getPrivateKey()
		if err != nil {
			return err
		}
		decryptKey, err := EccDecryptAction(ctx, content, privKey)
		if err != nil {
			return err
		}
		fmt.Printf("\033[1;32;40m Raw Password: %s\033[0m\n", decryptKey)
		os.Exit(codeOK)
	}

	return nil
}

func parseIO(input, output string) (*os.File, *os.File, error) {
	if input == "" || output == "" {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m invalid file path\033[0m\n")
	}
	in, err := os.Open(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m failed to open '%s': %v\033[0m\n", input, err)
		os.Exit(codeError) // TODO: replace by more gentle method
	}
	out, err := os.Create(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m failed to create '%s': %v\033[0m\n", output, err)
		os.Exit(codeError) // TODO: replace by more gentle method
	}
	cleanFn = append(cleanFn, func(code int) {
		out.Close()
		if code != codeOK { // remove file on error
			os.Remove(out.Name())
		}
	})
	return in, out, nil
}

func validateAesPswd(input string) error {
	if input == "" {
		return fmt.Errorf("invalid AES Key: no password")
	}
	if len(input) < 16 {
		return fmt.Errorf("at lease 16 characters")
	}
	return nil
}

func saveAesPswd(pswd string) (int64, error) {
	record := &model.KeyRecord{
		Password: pswd,
	}
	keyID, err := record.CreateRecord()
	if err != nil {
		return 0, err
	}
	return keyID, nil
}

func getAesKey() (string, error) {
	prompt := promptui.Prompt{
		Label:    "\033[1;36;40m Enter Raw Password\033[0m",
		Mask:     '*',
		Validate: validateAesPswd,
	}
	key, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return key, nil
}

func getPublicKey() ([]byte, error) {
	prompt := promptui.Prompt{
		Label: "Enter Public Key(64 bytes)",
		Mask:  '*',
	}
	key, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	keyByte, err := encrypt.DerivePublicKey(key)
	return keyByte, err
}

func getPrivateKey() ([]byte, error) {
	prompt := promptui.Prompt{
		Label: "Enter Private Key(32 bytes)",
		Mask:  '*',
	}
	key, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	keyByte, err := decrypt.DerivePrivateKey(key)
	return keyByte, err
}

func getEncryptedDecryptionKey() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter Encrypted Password",
	}
	encryptedKey, err := prompt.Run()
	if err != nil {
		return "", err
	}
	return encryptedKey, nil
}

func parseBooleanSelection(prompt string) (bool, error) {
	promptSelect := promptui.Select{
		Label: fmt.Sprintf("\033[1;36;40m %s\033[0m", prompt),
		Items: []string{YES, NO},
	}
	_, flag, err := promptSelect.Run()
	if err != nil {
		return false, err
	}
	if flag == NO {
		return false, nil
	}
	return true, nil
}

func getIpfsConfig() (*ipfs.IpfsConfig, error) {
	// **************** Host ******************
	prompt := promptui.Prompt{
		Label:   "\033[1;36;40m IPFS Host\033[0m",
		Default: "http://localhost",
	}
	host, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	// **************** API ******************
	prompt = promptui.Prompt{
		Label:   "\033[1;36;40m IPFS API Port\033[0m",
		Default: "5001",
	}
	api, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	apiPort, err := strconv.Atoi(api)
	if err != nil {
		return nil, err
	}

	// **************** Gateway ******************
	prompt = promptui.Prompt{
		Label:   "\033[1;36;40m IPFS Gateway Port\033[0m",
		Default: "8080",
	}
	gateway, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	gatewayPort, err := strconv.Atoi(gateway)
	if err != nil {
		return nil, err
	}

	// **************** IPFS Authorization ******************
	prompt = promptui.Prompt{
		Label: "\033[1;36;40m IPFS Peer ID (Authorization)\033[0m",
	}
	peerID, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	prompt = promptui.Prompt{
		Label: "\033[1;36;40m IPFS Public Key (Authorization)\033[0m",
		Mask:  '*',
	}
	pubkey, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	// **************** End ******************

	ipfsConfig := &ipfs.IpfsConfig{
		PeerID:      peerID,
		Pubkey:      pubkey,
		Host:        host,
		APIPort:     apiPort,
		GatewayPort: gatewayPort,
	}

	return ipfsConfig, nil
}

func IpfsUpload(cfg *ipfs.IpfsConfig, encryptFile string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	path, err := ipfs.Upload(ctx, cfg, encryptFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m failed to upload '%s': %v\033[0m\n", encryptFile, err)
		os.Exit(codeError)
	}
	fmt.Printf("\033[1;32;40m Upload successfully!\n Path: %s\033[0m\n", path)

}

func IpfsDownload(cfg *ipfs.IpfsConfig, locationUrl string, path string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	err := ipfs.Download(ctx, cfg, locationUrl, path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[1;31;40m failed to download '%s': %v\033[0m\n", path, err)
		os.Exit(codeError)
	}
	fmt.Printf("\033[1;32;40m Download successfully!\n Path: %s\033[0m\n", path)
}
