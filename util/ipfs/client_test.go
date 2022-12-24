package ipfs

import (
	"context"
	"fmt"
	"testing"
)

// id 12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH
// key CAESIEAhr0ClNGdjQTRCH3VJgaHMl8vi8wZS3pQ+XNddqVYe
// api http://localhost:5001
// README.md
// http://localhost:8080/ipfs/Qme1AwS6vUsAjdhfkNUGFdM49GQr6SXsg6xLZVqgsygacE
func TestUpload(t *testing.T) {
	cfg := IpfsConfig{
		PeerID:      "12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH",
		Pubkey:      "CAESIEAhr0ClNGdjQTRCH3VJgaHMl8vi8wZS3pQ",
		Host:        "http://localhost",
		APIPort:     5001,
		GatewayPort: 8080,
	}
	localFile := "../../README.md"
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	path, err := Upload(ctx, &cfg, localFile)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("path = %s\n", path)
}

func TestDownload(t *testing.T) {
	cfg := IpfsConfig{
		PeerID:      "12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH",
		Pubkey:      "CAESIEAhr0ClNGdjQTRCH3VJgaHMl8vi8wZS3pQ",
		Host:        "http://localhost",
		APIPort:     5001,
		GatewayPort: 8080,
	}
	locationUrl := "/ipfs/QmQfCC7AhVzb2By7DPpnCiNT4L9i3nFriciBmYFMufCY8v"
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	err := Download(ctx, &cfg, locationUrl, "./test.md")
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestStat(t *testing.T) {
	cfg := IpfsConfig{
		PeerID:      "12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH",
		Pubkey:      "CAESIEAhr0ClNGdjQTRCH3VJgaHMl8vi8wZS3pQ",
		Host:        "http://localhost",
		APIPort:     5001,
		GatewayPort: 8080,
	}
	cid := "Qme1AwS6vUsAjdhfkNUGFdM49GQr6SXsg6xLZVqgsygacD"
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	stat, err := Stat(ctx, &cfg, cid)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("%v\n", stat)
}
