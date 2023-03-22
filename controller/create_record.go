package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/types"
	"github.com/nextdotid/creator_suite/util"
	"golang.org/x/xerrors"
	"math/big"
	"net/http"
	"path/filepath"
	"strconv"
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
	fmt.Println(c.PostForm("network"))
	req.Network = types.Network(c.PostForm("network"))
	if !req.Network.IsValid() {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Cannot support the network right now"))
		return
	}
	req.ManagedContract = c.PostForm("managed_contract")
	req.PaymentTokenAddress = c.PostForm("payment_token_address")
	req.PaymentTokenAmount = c.PostForm("payment_token_amount")
	req.Password = c.PostForm("password")
	req.ContentName = c.PostForm("content_name")
	et, _ := strconv.ParseInt(c.PostForm("encryption_type"), 10, 64)
	req.EncryptionType = int8(et)
	req.Description = c.PostForm("description")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("get file error", err))
		return
	}

	filename := "/storage/" + file.Filename
	fmt.Printf("filename: %s", filename)
	if err = c.SaveUploadedFile(file, filename); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("fail to upload the file", err))
		return
	}
	fileExtension := filepath.Ext(filename)
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
		fileExtension, req.Network, req.ContentName, req.Description)
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
