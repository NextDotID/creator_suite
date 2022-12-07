package controller

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util"
	"github.com/nextdotid/creator_suite/util/encrypt"

	"golang.org/x/xerrors"
	"net/http"
)

type GetContentRequest struct {
	ContentID int64  `json:"content_id"`
	PublicKey string `json:"public_key"`
}

type GetContentResponse struct {
	EncryptedDecryptionKey string
}

func get_content(c *gin.Context) {
	req := GetContentRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	pub_key, err := util.StringToPublicKey(req.PublicKey)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error, publicKey invalid", err))
		return
	}

	content, err := model.FindContentByID(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	//TODO should use content.ManagedContract to match different contract
	is_paid, err := model.IsQualified(crypto.PubkeyToAddress(*pub_key).String(), uint64(content.AssetID))
	if is_paid == false || err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Can't find any payment record: %w", err))
		return
	}

	key, err := model.FindKeyPairByID(content.KeyID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	encrypt_key := encrypt.EncryptContentByPublicKey(key.PrivateKey, req.PublicKey)
	c.JSON(http.StatusOK, GetContentResponse{
		EncryptedDecryptionKey: encrypt_key,
	})
}
