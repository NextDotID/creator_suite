package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"golang.org/x/xerrors"
	"net/http"
	"strconv"
)

type ContentInfoRequest struct {
	ContentID int64 `json:"content_id"`
}

type ContentInfoResponse struct {
	Extension       string `json:"extension"`
	ManagedContract string `json:"managed_contract"`
	ContentName     string `json:"content_name"`
	Description     string `json:"description"`
	CreatorAddress  string `json:"creator_address"`
	KeyID           int64  `json:"key_id"`
	CreatedTime     string `json:"created_time"`
	UpdateTime      string `json:"update_time"`
}

func content_info(c *gin.Context) {
	req := ContentInfoRequest{}

	req.ContentID, _ = strconv.ParseInt(c.Query("content_id"), 10, 64)
	content, err := model.FindContentByID(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}

	c.JSON(http.StatusOK, ContentInfoResponse{
		Extension:       content.FileExtension,
		ManagedContract: content.ManagedContract,
		CreatorAddress:  content.CreatorAddress,
		ContentName:     content.ContentName,
		Description:     content.Description,
		KeyID:           content.KeyID,
		CreatedTime:     content.CreatedAt.String(),
		UpdateTime:      content.UpdatedAt.String(),
	})
}
