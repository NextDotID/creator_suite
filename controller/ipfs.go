package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/util/ipfs"
	"golang.org/x/xerrors"
)

type AliveRequest = ipfs.IpfsConfig

type AliveResponse struct {
	IsAlive bool   `json:"is_alive"`
	Message string `json:"message"`
}

func alive(c *gin.Context) {
	req := ipfs.IpfsConfig{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	isAlive, err := ipfs.Alive(&req)
	if err != nil {
		c.JSON(http.StatusOK, AliveResponse{isAlive, fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, AliveResponse{isAlive, ""})
}

type UploadRequest struct {
	LocalFile string          `json:"local_file"`
	Cfg       ipfs.IpfsConfig `json:"cfg"`
}

type UploadResponse struct {
	ContentLocateUrl string `json:"content_locate_url"`
}

func upload(c *gin.Context) {
	req := UploadRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	path, err := ipfs.Upload(ctx, &req.Cfg, req.LocalFile)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in IPFS Upload: %w", err))
		return
	}
	c.JSON(http.StatusOK, UploadResponse{ContentLocateUrl: path})
}
