package controller

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/nextdotid/creator_suite/util"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util/encrypt"
	log "github.com/sirupsen/logrus"

	"net/http"

	"golang.org/x/xerrors"
)

type GetContentRequest struct {
	ContentID           int64  `json:"content_id"`
	EncryptionPublicKey string `json:"encryption_public_key"`
	Signature           string `json:"signature"`
	SignaturePayload    string `json:"signature_payload"`
}

type GetContentResponse struct {
	EncryptedPassword string `json:"encrypted_password"`
	EncryptedResult   string `json:"encrypted_result"`
	EncryptionType    int8   `json:"encryption_type"`
	FileExtension     string `json:"file_extension"`
}

func get_content(c *gin.Context) {
	req := GetContentRequest{}

	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error", err))
		return
	}

	recoveredAddr, err := util.ValidSignatureAndGetTheAddress(req.SignaturePayload, req.Signature)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error, publicKey invalid: %w", err))
		return
	}

	err = validateAddressAndTimestamp(req.SignaturePayload, recoveredAddr)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Signature Verification error: %w", err))
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

	is_paid, err := model.IsQualified(content.ManagedContract, recoveredAddr, assetID)
	if !is_paid || err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Can't find any payment record: %w", err))
		return
	}

	var encrypted_result string
	var encrypted_password string
	if content.EncryptionType == model.ENCRYPTION_TYPE_AES {
		key, err := model.FindKeyRecordByID(content.KeyID)
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
			return
		}
		encrypted_password, err = encrypt.EncryptPasswordWithEncryptionPublicKey(req.EncryptionPublicKey, key.Password)
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Fail to give the encrypted password: %w", err))
			return
		}
		encrypted_result, err = getContent(pathJoin(STORAGE, strconv.FormatInt(content.ID, 10), content.ContentName+".enc"))
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Fail to give the encrypted content: %w", err))
			return
		}
	} else {
		encrypted_result, err = encrypt.EncryptFileWithEncryptionPublicKey(req.EncryptionPublicKey, pathJoin(STORAGE, strconv.FormatInt(content.ID, 10), content.ContentName))
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Fail to give the encrypted content: %w", err))
			return
		}
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

func validateAddressAndTimestamp(signaturePayload string, address string) error {
	reAcc := regexp.MustCompile("0x[a-fA-F0-9]{40}")
	matchedAcc := reAcc.FindString(signaturePayload)
	if matchedAcc != address {
		return xerrors.New("sign account doesn't match")
	}

	reTime := regexp.MustCompile(`timestamp:(\d+)`)
	matchedTime := reTime.FindStringSubmatch(signaturePayload)
	if len(matchedTime) < 2 {
		return xerrors.New("cannot get timestamp information from signature payload")
	}
	signTime, err := strconv.ParseInt(matchedTime[1], 10, 64)
	if err != nil || (time.Now().Unix()-signTime) > 60 {
		return xerrors.New("sign account doesn't match")
	}

	return nil
}
