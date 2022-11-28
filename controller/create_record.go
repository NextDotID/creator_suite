package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util/encrypt"
	"golang.org/x/xerrors"
	"net/http"
)

type CreateRecordRequest struct {
}

type CreateRecordResponse struct {
}

func create_record(c *gin.Context) {
	//generate key pair
	keypairs := encrypt.GenerateKeyPair()
	key_record := &model.KeyPair{}
	key_record.PublicKey = keypairs[0]
	key_record.PrivateKey = keypairs[1]
	keyID, err := key_record.CreateRecord()
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}
	// encrypted content && upload to ipfs
	location := "test1"

	content := &model.Content{}
	content.LocationUrl = location
	content.KeyID = keyID
	err = content.CreateRecord()
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	c.JSON(http.StatusOK, CreateRecordResponse{})
}
