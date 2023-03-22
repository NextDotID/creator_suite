package controller

import (
	"github.com/nextdotid/creator_suite/types"
	"github.com/nextdotid/creator_suite/util"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"golang.org/x/xerrors"
)

type CreateRecordRequest struct {
	ManagedContract     string        `json:"managed_contract"`
	Network             types.Network `json:"network"`
	PaymentTokenAddress string        `json:"payment_token_address"`
	PaymentTokenAmount  string        `json:"payment_token_amount"`
	Password            string        `json:"password"`
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
		//fmt.Println(req)
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error", err))
		return
	}
	if !req.Network.IsValid() {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Cannot support the network right now"))
		return
	}

	var keyID int64
	if req.EncryptionType == model.ENCRYPTION_TYPE_AES {
		record := &model.KeyRecord{
			Password: req.Password,
		}
		var err error
		keyID, err = record.CreateRecord()
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
			return
		}
	}

	tokenAmount := util.ToWei(req.PaymentTokenAmount, 18)
	if tokenAmount == big.NewInt(0) {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("token amount invalid"))
		return
	}

	content, err := model.CreateRecord(req.ManagedContract, keyID, req.EncryptionType,
		req.FileExtension, req.Network, req.ContentName, req.Description)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	//err = model.CreateAsset(content.ID, req.ManagedContract, req.PaymentTokenAddress, tokenAmount, req.Network)
	//if err != nil {
	//	updateErr := content.UpdateToInvalidStatus(content.ID)
	//	if updateErr != nil {
	//		log.Errorf("update content record err:%v", updateErr)
	//	}
	//	errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Create an asset in Contract error: %v", err))
	//	return
	//}

	c.JSON(http.StatusOK, CreateRecordResponse{
		ContentID: content.ID,
	})
}
