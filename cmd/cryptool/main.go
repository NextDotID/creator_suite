package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/nextdotid/creator_suite/common"
	"github.com/nextdotid/creator_suite/util"
	"github.com/nextdotid/creator_suite/util/dare"
	"github.com/nextdotid/creator_suite/util/decrypt"
	"github.com/nextdotid/creator_suite/util/encrypt"
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
)

var cryptoFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "in",
		Value: "",
		Usage: "input filepath",
	},
	cli.StringFlag{
		Name:  "out",
		Value: "",
		Usage: "output filepath",
	},
	cli.StringFlag{
		Name:  "algorithm,alg",
		Value: "",
		Usage: "cryto algorithm, currently support AES and ECC",
	},
}

var CryptoolCommands = []cli.Command{
	{
		Name:    common.ENCRYPT,
		Aliases: []string{"encrypt", "en"},
		Usage:   "encrypt a file",
		Flags:   cryptoFlags,
		Action: func(ctx *cli.Context) error {
			return commandAction(ctx, common.ENCRYPT)
		},
	},
	{
		Name:    common.DECRYPT,
		Aliases: []string{"decrypt", "de"},
		Usage:   "decrypt a file",
		Flags:   cryptoFlags,
		Action: func(ctx *cli.Context) error {
			return commandAction(ctx, common.DECRYPT)
		},
	},
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

func main() {
	cleanChan := make(chan int, 1)
	go func() {
		code := <-cleanChan
		for _, f := range cleanFn {
			f(code)
		}
		os.Exit(code)
	}()
	err := InitCmdTool()
	if err != nil {
		fmt.Printf("err = %v\n", err)
		panic(err)
	}
}

func chooseAlgorithm(ctx *cli.Context) (string, error) {
	algorithm := ""
	items := []string{common.AES, common.ECC}
	if ctx.String("algorithm") != "" {
		algorithm = ctx.String("algorithm")
	} else {
		prompt := promptui.Select{
			Label: "Select Algorithm",
			Items: items,
		}
		_, alg, err := prompt.Run()

		if err != nil {
			return "", err
		}
		algorithm = alg
	}
	if algorithm == "" {
		return "", fmt.Errorf("no crypto algorithm selected")
	}

	if !util.StringsIn(items, algorithm) {
		return "", fmt.Errorf("unsupported crypto algorithm[%s]", algorithm)
	}
	return algorithm, nil
}

func getPassword(ctx *cli.Context) (string, error) {
	var key string
	var err error
	prompt := promptui.Prompt{
		Label: "Enter Password",
		Mask:  '*',
	}
	key, err = prompt.Run()
	if err != nil {
		return "", err
	}
	// if key == "" {
	// 	return "", fmt.Errorf("failed to AES Key: No password")
	// }
	// if len(key) != 128 {
	// 	return "", fmt.Errorf("AES Key must be equal to 128 bit")
	// }
	return key, nil
}

func parseIO(ctx *cli.Context) (*os.File, *os.File) {
	if ctx.String("in") == "" || ctx.String("out") == "" {
		return nil, nil
	}
	in, err := os.Open(ctx.String("in"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open '%s': %v\n", ctx.String("in"), err)
		os.Exit(codeError) // TODO: replace by more gentle method
	}
	out, err := os.Create(ctx.String("out"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create '%s': %v\n", ctx.String("out"), err)
		os.Exit(codeError) // TODO: replace by more gentle method
	}
	cleanFn = append(cleanFn, func(code int) {
		out.Close()
		if code != codeOK { // remove file on error
			os.Remove(out.Name())
		}
	})
	return in, out
}

func commandAction(ctx *cli.Context, action string) error {
	algorithm, err := chooseAlgorithm(ctx)
	if err != nil {
		return err
	}
	in, out := parseIO(ctx)
	pswd, err := getPassword(ctx)
	if err != nil {
		return err
	}
	switch algorithm {
	default:
		err = fmt.Errorf("unsupported crypto algorithm[%s]", algorithm)
	case common.AES:
		if action == common.ENCRYPT {
			err = AesEncryptAction(ctx, pswd, in, out)
			if err != nil {
				return err
			}
		} else if action == common.DECRYPT {
			err = AesDecryptAction(ctx, pswd, in, out)
			if err != nil {
				return err
			}
		}
	case common.ECC:
		content := ""
		if action == common.ENCRYPT {
			content, err = EccEncryptAction(ctx)
			fmt.Printf("content: %v\n", content)
		} else if action == common.DECRYPT {
			content, err = EccDecryptAction(ctx)
			fmt.Printf("content: %v\n", content)
		}
	}
	return err
}

func EccEncryptAction(ctx *cli.Context) (string, error) {
	return "EccEncryptAction", nil
}

func EccDecryptAction(ctx *cli.Context) (string, error) {
	return "EccDecryptAction", nil
}

func AesEncryptAction(ctx *cli.Context, aesKey string, src, dst *os.File) error {
	key, err := encrypt.DeriveKey([]byte(aesKey), src, dst)
	if err != nil {
		return err
	}
	cfg := dare.Config{Key: key}
	if _, err := encrypt.AesEncrypt(src, dst, cfg); err != nil {
		return err
	}
	return nil
}

func AesDecryptAction(ctx *cli.Context, aesKey string, src, dst *os.File) error {
	key, err := decrypt.DeriveKey([]byte(aesKey), src, dst)
	if err != nil {
		return err
	}
	cfg := dare.Config{Key: key}
	if _, err := decrypt.AesDecrypt(src, dst, cfg); err != nil {
		return err
	}
	return nil
}
