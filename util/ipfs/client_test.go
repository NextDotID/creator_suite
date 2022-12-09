package ipfs

import (
	"context"
	"fmt"
	"testing"
)

// id 12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH
// key CAESIEAhr0ClNGdjQTRCH3VJgaHMl8vi8wZS3pQ+XNddqVYe
// api http://localhost:5001
// test-nextdotid.png
// http://localhost:8080/ipfs/QmQfCC7AhVzb2By7DPpnCiNT4L9i3nFriciBmYFMufCY8v
func TestUpload(t *testing.T) {
	cfg := IpfsConfig{
		NodeID:      "12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH",
		Pubkey:      "CAESIEAhr0ClNGdjQTRCH3VJgaHMl8vi8wZS3pQ",
		Host:        "http://localhost",
		APIPort:     5001,
		GatewayPort: 8080,
	}
	localFile := "../../test-nextdotid.png"
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
		NodeID:      "12D3KooWE8i3z8DPU2tQBedoMez7dfo9nNrNcCH4i6vRLb6ZrYzH",
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
	err := Download(ctx, &cfg, locationUrl)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
