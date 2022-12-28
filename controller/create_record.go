package controller

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"golang.org/x/xerrors"
)

type CreateRecordRequest struct {
	ContentLocateUrl    string `json:"content_locate_url"`
	ManagedContract     string `json:"managed_contract"`
	PaymentTokenAddress string `json:"payment_token_address"`
	PaymentTokenAmount  int64  `json:"payment_token_amount"`
	KeyID               int64  `json:"key_id"`
	EncryptionType      int8   `json:"encryption_type"`
	FileExtension       string `json:"file_extension"`
}

type CreateRecordResponse struct {
}

func create_record(c *gin.Context) {
	req := CreateRecordRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	kr, err := model.FindKeyRecordByID(req.KeyID)
	if err != nil || kr == nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error, cannot find encryption key"))
		return
	}

	content, err := model.CreateRecord(req.ContentLocateUrl, req.ManagedContract, req.KeyID, req.EncryptionType, req.FileExtension)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	// create asset in contract, TODO should be multiple contract options
	err = model.CreateAsset(content.ID, req.ManagedContract, req.PaymentTokenAddress, req.PaymentTokenAmount)
	if err != nil {
		err = content.UpdateToInvalidStatus(content.ID)
		if err != nil {
			log.Errorf("update content record err:%v", err)
		}
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Create an asset in Contract error: %w", err))
		return
	}

	c.JSON(http.StatusOK, CreateRecordResponse{})
}
