package controller

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// STORAGE which defines in docker-compose
const STORAGE = "./storage"

type ShowContentRequest struct {
	ContentID int64 `json:"content_id"`
}

type ShowContentResponse struct {
	ContentID      int64  `json:"content_id"`
	FileName       string `json:"file_name"`
	PreviewImage   []byte `json:"preview_image"`
	Description    string `json:"description"`
	EncryptionType int8   `json:"encryption_type"`
	FileExtension  string `json:"file_extension"`
	FileSize       string `json:"file_size"`

	AssetID             uint64 `json:"asset_id"`
	Network             string `json:"network"`
	PaymentTokenAddress string `json:"payment_token_address"`
	PaymentTokenAmount  int64  `json:"payment_token_amount"`
}

func show_content(c *gin.Context) {
	req := ShowContentRequest{}
	req.ContentID, _ = strconv.ParseInt(c.Query("content_id"), 0, 64)

	content, err := model.FindContentByID(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	// TODO: query token_amount from contract
	assetID, err := model.GetAssetID(content.CreatorAddress, uint64(content.ID))
	log.Infof("get assetID: %d", assetID)

	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in read contract record: %w", err))
		return
	}

	resp := ShowContentResponse{
		ContentID:      content.ID,
		PreviewImage:   nil,
		Description:    content.Description,
		EncryptionType: content.EncryptionType,
		FileExtension:  content.FileExtension,

		AssetID:             assetID,
		Network:             content.Network,
		PaymentTokenAddress: "",
		PaymentTokenAmount:  1,
	}

	// STORAGE: Get File Information
	list, err := ioutil.ReadDir(pathJoin(STORAGE, strconv.FormatInt(req.ContentID, 10)))
	if err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
		return
	}
	for _, item := range list {
		if !item.IsDir() {
			if item.Name() == "preview.png" || item.Name() == "preview.jpg" || item.Name() == "preview.jpeg" {
				resp.PreviewImage = nil
			} else {
				resp.FileName = strings.TrimRight(item.Name(), ".enc")
				resp.FileSize = formatFileSize(item.Size())
			}
		}
	}
	c.JSON(http.StatusOK, resp)
}
