package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util/encrypt"
	"golang.org/x/xerrors"
	"net/http"
)

type GetContentRequest struct {
	ContentID int64
	publicKey string
}

type GetContentResponse struct {
	EncryptedDecryptionKey string
}

func get_content(c *gin.Context) {
	req := GetContentRequest{}
	if err := c.BindQuery(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	content, err := model.FindContentByID(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	//TODO query isQualified() function from contract

	key, err := model.FindKeyPairByID(content.KeyID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	encrypt_key := encrypt.EncryptContentByPublicKey(key.PrivateKey, req.publicKey)
	c.JSON(http.StatusOK, GetContentResponse{
		EncryptedDecryptionKey: encrypt_key,
	})
}
