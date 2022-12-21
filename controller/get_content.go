package controller

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util"
	"github.com/nextdotid/creator_suite/util/encrypt"
	log "github.com/sirupsen/logrus"
	"strconv"

	"net/http"

	"golang.org/x/xerrors"
)

type GetContentRequest struct {
	ContentID int64  `json:"content_id"`
	PublicKey string `json:"public_key"`
}

type GetContentResponse struct {
	EncryptedDecryptionKey string `json:"encrypted_decryption_key"`
	LocationUrl            string `json:"location_url"`
}

func get_content(c *gin.Context) {
	req := GetContentRequest{}
	req.PublicKey = c.Query("public_key")
	req.ContentID, _ = strconv.ParseInt(c.Query("content_id"), 0, 64)

	pub_key, err := util.StringToPublicKey(req.PublicKey)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error, publicKey invalid: %w", err))
		return
	}

	content, err := model.FindContentByID(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	assetID, err := model.GetAssetID(content.CreatorAddress, uint64(content.ID))
	log.Infof("get assetID: %d", assetID)

	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in read contract record: %w", err))
		return
	}

	//TODO should use content.ManagedContract to match different contract
	is_paid, err := model.IsQualified(crypto.PubkeyToAddress(*pub_key).String(), assetID)
	if !is_paid || err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Can't find any payment record: %w", err))
		return
	}

	key, err := model.FindKeyRecordByID(content.KeyID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	encrypt_key, err := encrypt.EncryptContentByPublicKey(key.Password, req.PublicKey)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Can't encrypt content by public_key: %w", err))
		return
	}

	c.JSON(http.StatusOK, GetContentResponse{
		EncryptedDecryptionKey: encrypt_key,
		LocationUrl:            content.LocationUrl,
	})
}
