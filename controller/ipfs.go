package controller

import (
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
	req := AliveRequest{}
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
