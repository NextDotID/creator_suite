package controller

import (
	"net/http"
	"strconv"

	"github.com/nextdotid/creator_suite/types"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"golang.org/x/xerrors"
)

type CreateRecordRequest struct {
	ContentLocateUrl    string        `json:"content_locate_url"`
	ManagedContract     string        `json:"managed_contract"`
	Network             types.Network `json:"network"`
	PaymentTokenAddress string        `json:"payment_token_address"`
	PaymentTokenAmount  int64         `json:"payment_token_amount"`
	KeyID               int64         `json:"key_id"`
	ContentName         string        `json:"content_name"`
	EncryptionType      int8          `json:"encryption_type"`
	FileExtension       string        `json:"file_extension"`
	Description         string        `json:"description"`
}

type CreateRecordResponse struct {
	ContentID int64 `json:"content_id"`
}

func create_record(c *gin.Context) {
	req := CreateRecordRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}
	if !req.Network.IsValid() {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Cannot support the network right now"))
		return
	}

	if req.EncryptionType == model.ENCRYPTION_TYPE_AES {
		kr, err := model.FindKeyRecordByID(req.KeyID)
		if err != nil || kr == nil {
			errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error, cannot find encryption key"))
			return
		}
	}

	content, err := model.CreateRecord(req.ContentLocateUrl, req.ManagedContract, req.KeyID, req.EncryptionType,
		req.FileExtension, req.Network, req.ContentName, req.Description)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	if req.EncryptionType == model.ENCRYPTION_TYPE_ECC {
		err = content.UpdateLocationUrl(pathJoin(STORAGE, strconv.FormatInt(content.ID, 10), req.ContentName))
		if err != nil {
			log.Errorf("update content_url err: %v", err)
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Update error: %v", err))
			return
		}
	}

	// create asset in contract, TODO should be multiple contract options
	err = model.CreateAsset(content.ID, req.ManagedContract, req.PaymentTokenAddress, req.PaymentTokenAmount)
	if err != nil {
		updateErr := content.UpdateToInvalidStatus(content.ID)
		if updateErr != nil {
			log.Errorf("update content record err:%v", updateErr)
		}
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Create an asset in Contract error: %v", err))
		return
	}

	c.JSON(http.StatusOK, CreateRecordResponse{
		ContentID: content.ID,
	})
}
