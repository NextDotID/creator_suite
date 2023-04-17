package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type Folder struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	CreatedTime string `json:"created_time"`
	UpdateTime  string `json:"update_time"`
	ContentID   int64  `json:"content_id"`
	Files       []File `json:"files"`
}

type File struct {
	Name            string `json:"name"`
	Size            string `json:"size"`
	Extension       string `json:"extension"`
	Path            string `json:"path"`
	ManagedContract string `json:"managed_contract"`
	ContentName     string `json:"content_name"`
	Description     string `json:"description"`
	CreatorAddress  string `json:"creator_address"`
	KeyID           int64  `json:"key_id"`
	CreatedTime     string `json:"created_time"`
	UpdateTime      string `json:"update_time"`
}

type ListResponse struct {
	Folders []Folder `json:"folders"`
}

func formatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func Ext(name string) string {
	suffix := filepath.Ext(name)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return "other"
}

func pathJoin(basePath string, elem ...string) string {
	if strings.HasPrefix(basePath, "./") {
		base := strings.TrimSuffix(basePath, "/")
		return base + "/" + filepath.Join(elem...)
	}
	return filepath.Join(basePath, filepath.Join(elem...))
}

func list(c *gin.Context) {
	const FILE_PATH = "/storage"
	folders := make([]Folder, 0)

	// list content table
	contents, err := model.ListContent()
	if err != nil {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
		return
	}
	contentMap := make(map[int64]model.Content)
	for _, c := range contents {
		contentMap[c.ID] = c
	}

	list, err := ioutil.ReadDir(FILE_PATH)
	if err != nil {
		log.Infof("I/O error: %v", err)
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
		return
	}
	for _, item := range list {
		if item.IsDir() {
			contentID, err := strconv.ParseInt(item.Name(), 10, 32)
			if err != nil {
				log.Warnf("invalid contentID: %d err: %v", contentID, err)
				continue
			}
			content := contentMap[contentID]
			if &content == nil {
				log.Warnf("fail to get content info, contentID: %d", contentID)
				continue
			}
			folder := Folder{
				Name:        item.Name(),
				Path:        pathJoin(FILE_PATH, item.Name()),
				ContentID:   contentID,
				CreatedTime: util.Datetime2DateString(item.ModTime()),
				UpdateTime:  util.Datetime2DateString(item.ModTime()),
			}
			files := make([]File, 0)
			f, err := ioutil.ReadDir(folder.Path)
			if err != nil {
				errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
				return
			}
			for _, item := range f {
				if !item.IsDir() {
					files = append(files, File{
						Name:            item.Name(),
						Size:            formatFileSize(item.Size()),
						Extension:       content.FileExtension,
						ManagedContract: content.ManagedContract,
						CreatorAddress:  content.CreatorAddress,
						ContentName:     content.ContentName,
						Description:     content.Description,
						KeyID:           content.KeyID,
						Path:            filepath.Join(folder.Path, item.Name()),
						CreatedTime:     util.Datetime2DateString(item.ModTime()),
						UpdateTime:      util.Datetime2DateString(item.ModTime()),
					})
				}
			}
			folder.Files = files
			folders = append(folders, folder)
		}
	}

	c.JSON(http.StatusOK, ListResponse{
		Folders: folders,
	})
}
