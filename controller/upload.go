package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/xerrors"
	"net/http"
	"path/filepath"
)

type UploadFileResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func upload_file(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("get file error", err))
		return
	}

	filename := filepath.Dir("/storage/" + file.Filename)
	fmt.Printf("filename: %s", filename)

	if err = c.SaveUploadedFile(file, filename); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("fail to upload the file", err))
		return
	}

	c.JSON(http.StatusOK, UploadFileResponse{
		Name:  name,
		Email: email,
	})

}
