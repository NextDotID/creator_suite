package controller

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io/ioutil"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util"
	"github.com/nextdotid/creator_suite/util/encrypt"
	log "github.com/sirupsen/logrus"

	"net/http"

	"golang.org/x/xerrors"
)

type GetContentRequest struct {
	ContentID int64  `json:"content_id"`
	PublicKey string `json:"public_key"`
}

type GetContentResponse struct {
	EncryptedPassword string `json:"encrypted_password"`
	EncryptedResult   string `json:"encrypted_result"`
	EncryptionType    int8   `json:"encryption_type"`
	FileExtension     string `json:"file_extension"`
}

func get_content(c *gin.Context) {
	req := GetContentRequest{}
	req.PublicKey = c.Query("public_key")
	req.ContentID, _ = strconv.ParseInt(c.Query("content_id"), 10, 64)

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

	assetID, err := model.GetAssetID(content.ManagedContract, content.CreatorAddress, uint64(content.ID))
	log.Infof("get assetID: %d", assetID)

	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in read contract record: %w", err))
		return
	}

	is_paid, err := model.IsQualified(content.ManagedContract, crypto.PubkeyToAddress(*pub_key).String(), assetID)
	if !is_paid || err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Can't find any payment record: %w", err))
		return
	}

	key, err := model.FindKeyRecordByID(content.KeyID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	var encrypted_result string
	var encrypted_password string
	if content.EncryptionType == model.ENCRYPTION_TYPE_AES {
		encrypted_password, err = encrypt.EncryptPasswordByPublicKey(key.Password, req.PublicKey)
		encrypted_result, err = getContent(pathJoin(STORAGE, strconv.FormatInt(content.ID, 10), content.ContentName+".enc"))
	} else {
		encrypted_result, err = encrypt.EncryptContentByPublicKey(pathJoin(STORAGE, strconv.FormatInt(content.ID, 10), content.ContentName), req.PublicKey)
	}
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Can't encrypt content by public_key: %w", err))
		return
	}

	c.JSON(http.StatusOK, GetContentResponse{
		EncryptedPassword: encrypted_password,
		EncryptedResult:   encrypted_result,
		EncryptionType:    content.EncryptionType,
		FileExtension:     content.FileExtension,
	})
}

func getContent(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("invalid public key", err)
	}
	return hexutil.Encode(bytes), nil
}
