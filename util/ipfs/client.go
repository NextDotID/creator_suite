package ipfs

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	ipfsfiles "github.com/ipfs/go-ipfs-files"
	ipfshttpcli "github.com/ipfs/go-ipfs-http-client"
	caopts "github.com/ipfs/interface-go-ipfs-core/options"

	ipfspath "github.com/ipfs/interface-go-ipfs-core/path"
)

// IpfsFile IPFS file infomation
type IpfsFile struct {
	Name string
	Cid  string // the CID of the root object of the path
	Path string // /ipfs/xxxxx path from the provided CID
	Size int64
}

func (fs IpfsFile) String() string {
	return fmt.Sprintf("Name %v | Cid %v, Path: %v | Size: %v", fs.Name, fs.Cid, fs.Path, fs.Size)
}

// IpfsConfig IPFS Authorization Config
type IpfsConfig struct {
	PeerID      string `json:"peer_id"`
	Pubkey      string `json:"pubkey"`
	Host        string `json:"host"`
	APIPort     int    `json:"api_port"`
	GatewayPort int    `json:"gateway_port"`
}

type IpfsStat struct {
	Hash string
	Type string
	Size int64 // unixfs size
}

/**
 * @description:
 * @param {*IpfsConfig} cfg IPFS Authorization Config
 * @param {string} path LocalFile path
 * @return {
	{string} LocationUrl: object of the path in IPFS
 }
*/
func Upload(ctx context.Context, cfg *IpfsConfig, path string) (string, error) {
	httpCli := &http.Client{}
	cli, err := ipfshttpcli.NewURLApiWithClient(apiUrl(cfg), httpCli)
	if err != nil {
		return "", err
	}
	cli.Headers.Add("Authorization", auth(cfg.PeerID, cfg.Pubkey))

	stat, err := os.Lstat(path)
	if err != nil {
		return "", err
	}
	file, err := ipfsfiles.NewSerialFile(path, false, stat)
	if err != nil {
		return "", nil
	}

	var res ipfspath.Resolved
	var ipfsfile IpfsFile
	start := time.Now().Unix()
	res, err = cli.Unixfs().Add(
		ctx,
		file,
		caopts.Unixfs.Pin(false),
		caopts.Unixfs.Progress(true),
	)
	if err != nil {
		return "", err
	}
	ipfsfile.Name = stat.Name()
	ipfsfile.Size = stat.Size()
	ipfsfile.Cid = res.Cid().String()
	ipfsfile.Path = parseGatewayPath(cfg, ipfsfile.Cid)

	verbose := false // TODO: add verbose into configuration
	if verbose {
		log.Info(ipfsfile.String())
		log.Info(fmt.Sprintf("cid=%v time cost: %v seconds...", res.Cid().String(), time.Now().Unix()-start))
	}

	// call cp
	var cpRes interface{}
	cli.Request("files/cp",
		parseIpfsPath(ipfsfile.Cid),
		parseName(ipfsfile.Name),
		"stream-channels=true").Exec(ctx, &cpRes)
	return ipfsfile.Path, nil
}

func Stat(ctx context.Context, cfg *IpfsConfig, cid string) (*IpfsStat, error) {
	httpCli := &http.Client{Timeout: 5 * time.Second}
	cli, err := ipfshttpcli.NewURLApiWithClient(apiUrl(cfg), httpCli)
	if err != nil {
		return nil, err
	}
	cli.Headers.Add("Authorization", auth(cfg.PeerID, cfg.Pubkey))

	var stat IpfsStat
	err = cli.Request("files/stat", ParseIpfsPath(cid)).Exec(ctx, &stat)
	if err != nil {
		return nil, err
	}
	return &stat, nil
}

func Ls(ctx context.Context, cfg *IpfsConfig, locationUrl string) error {
	httpCli := &http.Client{}
	cli, err := ipfshttpcli.NewURLApiWithClient(apiUrl(cfg), httpCli)
	if err != nil {
		return err
	}
	res, err := cli.Unixfs().Ls(ctx, ipfspath.New(locationUrl))
	if err != nil {
		return err
	}
	entry := <-res
	if entry.Err != nil {
		return entry.Err
	}
	fmt.Printf("%v\n", entry)
	return nil
}

func Download(ctx context.Context, cfg *IpfsConfig, locationUrl string, path string) error {
	httpCli := &http.Client{}
	cli, err := ipfshttpcli.NewURLApiWithClient(apiUrl(cfg), httpCli)
	if err != nil {
		return err
	}
	cli.Headers.Add("Authorization", auth(cfg.PeerID, cfg.Pubkey))
	res, err := cli.Unixfs().Get(ctx, ipfspath.New(locationUrl))
	if err != nil {
		return err
	}
	ipfsfiles.WriteTo(res, path)
	return nil
}

func auth(id, pubkey string) string {
	auth := id + ":" + pubkey
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func apiUrl(cfg *IpfsConfig) string {
	return cfg.Host + ":" + strconv.Itoa(cfg.APIPort)
}

func ParseGatewayPath(cfg *IpfsConfig, cid string) string {
	return cfg.Host + ":" + strconv.Itoa(cfg.GatewayPort) + filepath.Join("/ipfs", cid)
}

func ParseIpfsPath(cid string) string {
	return filepath.Join("/ipfs", cid)
}

func ParseName(name string) string {
	return filepath.Join("/", name)
}

func ParseCid(path string) string {
	strs := strings.Split(path, "/")
	return strs[len(strs)-1]
}
