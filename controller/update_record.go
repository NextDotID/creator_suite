package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"golang.org/x/xerrors"
	"net/http"
)

type UpdateRecordRequest struct {
	ContentID int64 `json:"content_id"`
	//AssetID   int64 `json:"stat"`
}

type UpdateRecordResponse struct {
	IsOk int
}

func update_record(c *gin.Context) {
	req := UpdateRecordRequest{}
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(req)
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error", err))
		return
	}

	content, err := model.FindContentByID(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		c.JSON(http.StatusOK, UpdateRecordResponse{})
		return
	}
	err = content.UpdateToInvalidStatus(req.ContentID)
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		c.JSON(http.StatusOK, UpdateRecordResponse{})
		return
	}
	c.JSON(http.StatusOK, UpdateRecordResponse{
		IsOk: 1,
	})
}
