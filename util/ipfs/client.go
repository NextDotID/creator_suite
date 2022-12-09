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
	NodeID      string `json:"node_id"`
	Pubkey      string `json:"pubkey"`
	Host        string `json:"host"`
	APIPort     int    `json:"api_port"`
	GatewayPort int    `json:"gateway_port"`
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
	cli.Headers.Add("Authorization", auth(cfg.NodeID, cfg.Pubkey))

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
	ipfsfile.Path = parsePath(cfg, ipfsfile.Cid)

	log.Info(ipfsfile.String())
	log.Info(fmt.Sprintf("cid=%v time cost: %v seconds...", res.Cid().String(), time.Now().Unix()-start))

	// call cp
	var cpRes interface{}
	err = cli.Request("files/cp", fmt.Sprintf("/ipfs/%s", ipfsfile.Cid), "/test-nextdotid.png", "stream-channels=true").Exec(ctx, &cpRes)
	if err != nil {
		return ipfsfile.Path, nil
	}
	return ipfsfile.Path, nil
}

func Download(ctx context.Context, cfg *IpfsConfig, locationUrl string) error {
	httpCli := &http.Client{}
	cli, err := ipfshttpcli.NewURLApiWithClient(apiUrl(cfg), httpCli)
	if err != nil {
		return err
	}
	cli.Headers.Add("Authorization", auth(cfg.NodeID, cfg.Pubkey))
	res, err := cli.Unixfs().Get(ctx, ipfspath.New(locationUrl))
	if err != nil {
		return err
	}
	ipfsfiles.WriteTo(res, "./test.md")
	return nil
}

func auth(id, pubkey string) string {
	auth := id + ":" + pubkey
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func apiUrl(cfg *IpfsConfig) string {
	return cfg.Host + ":" + strconv.Itoa(cfg.APIPort)
}

func parsePath(cfg *IpfsConfig, cid string) string {
	return cfg.Host + ":" + strconv.Itoa(cfg.GatewayPort) + filepath.Join("/ipfs", cid)
}
