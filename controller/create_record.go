package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/types"
	"github.com/nextdotid/creator_suite/util/dare"
	"github.com/nextdotid/creator_suite/util/encrypt"
	"golang.org/x/xerrors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type CreateRecordRequest struct {
	ManagedContract string        `json:"managed_contract"`
	Network         types.Network `json:"network"`
	Password        string        `json:"password"`
	EncryptionType  int8          `json:"encryption_type"`
	FileExtension   string        `json:"file_extension"`
	Description     string        `json:"description"`
	CreatorAddress  string        `json:"creator_address"`
}

type CreateRecordResponse struct {
	ContentID int64 `json:"content_id"`
}

func create_record(c *gin.Context) {
	req := CreateRecordRequest{}
	req.Network = types.Network(c.PostForm("network"))
	if !req.Network.IsValid() {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Cannot support the network right now"))
		return
	}
	req.ManagedContract = c.PostForm("managed_contract")
	req.Password = c.PostForm("password")
	et, _ := strconv.ParseInt(c.PostForm("encryption_type"), 10, 64)
	req.EncryptionType = int8(et)
	req.Description = c.PostForm("description")
	req.CreatorAddress = c.PostForm("creator_address")

	file, err := c.FormFile("file")
	if err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("get file error", err))
		return
	}
	fileExtension := strings.Trim(filepath.Ext(file.Filename), ".")

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

	content, err := model.CreateRecord(req.ManagedContract, keyID, req.EncryptionType,
		fileExtension, req.Network, file.Filename, req.Description, req.CreatorAddress)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	filePath := "/storage/" + strconv.FormatInt(content.ID, 10) + "/" + file.Filename
	if err = os.Mkdir("/storage/"+strconv.FormatInt(content.ID, 10), os.ModePerm); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("fail to create new folder, err: %w", err))
		return
	}
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("fail to upload the file, err:%w", err))
		return
	}

	// generate encrypted file
	if req.EncryptionType == model.ENCRYPTION_TYPE_AES {
		src, _ := os.Open(filePath)
		dst, _ := os.Create(filePath + ".enc")
		key, err := encrypt.DeriveKey([]byte(req.Password), src, dst)
		if err != nil {
			errorResp(c, http.StatusBadRequest, xerrors.Errorf("fail to DeriveKey to encrypt", err))
			return
		}
		cfg := dare.Config{Key: key}
		_, err = encrypt.AesEncrypt(src, dst, cfg)
		if err != nil {
			errorResp(c, http.StatusBadRequest, xerrors.Errorf("fail to encrypt the file", err))
			return
		}
	}

	c.JSON(http.StatusOK, CreateRecordResponse{
		ContentID: content.ID,
	})
}
