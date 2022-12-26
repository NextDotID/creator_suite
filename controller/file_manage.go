package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nextdotid/creator_suite/model"
	"github.com/nextdotid/creator_suite/util"
	"github.com/nextdotid/creator_suite/util/dare"
	"github.com/nextdotid/creator_suite/util/decrypt"
	"github.com/nextdotid/creator_suite/util/encrypt"
	"github.com/nextdotid/creator_suite/util/ipfs"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type Folder struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Path        string `json:"path"`
	CreatedTime string `json:"created_time"`
	UpdateTime  string `json:"update_time"`
	ContentID   int64  `json:"content_id"`
}

type File struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Size      string `json:"size"`
	Extension string `json:"extension"`
	Path      string `json:"path"`

	ContentID       int64  `json:"content_id"`
	ManagedContract string `json:"managed_contract"`
	AssetID         int64  `json:"asset_id"`
	KeyID           int64  `json:"key_id"`
	LocationUrl     string `json:"location_url"`
	CreatedTime     string `json:"created_time"`
	UpdateTime      string `json:"update_time"`
}

type ListRequest struct {
	Path string          `json:"path"` // /storage/
	Cfg  ipfs.IpfsConfig `json:"cfg"`
}

type ListResponse struct {
	Folders []Folder `json:"folders"`
	Files   []File   `json:"files"`
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

func list(c *gin.Context) {
	req := ListRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	folders := make([]Folder, 0)
	files := make([]File, 0)
	list, err := ioutil.ReadDir(req.Path)
	if err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
		return
	}

	for _, item := range list {
		if item.IsDir() {
			contentID, err := strconv.ParseInt(item.Name(), 32, 10)
			if err != nil {
				continue
			}
			folders = append(folders, Folder{
				Name:        item.Name(),
				Type:        "dirs",
				Path:        filepath.Join(req.Path, item.Name()),
				ContentID:   contentID,
				CreatedTime: util.Datetime2DateString(item.ModTime()),
				UpdateTime:  util.Datetime2DateString(item.ModTime()),
			})
		} else {
			files = append(files, File{
				Name:        item.Name(),
				Type:        "localfile",
				Size:        formatFileSize(item.Size()),
				Extension:   Ext(item.Name()),
				Path:        filepath.Join(req.Path, item.Name()),
				CreatedTime: util.Datetime2DateString(item.ModTime()),
				UpdateTime:  util.Datetime2DateString(item.ModTime()),
			})
		}
	}

	// for _, folder := range folders {
	// 	f, err := ioutil.ReadDir(folder.Path)
	// 	if err != nil {
	// 		errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
	// 		return
	// 	}
	// 	for _, item := range f {
	// 		if !item.IsDir() {
	// 			files = append(files, File{
	// 				Name:        item.Name(),
	// 				Type:        "localfile",
	// 				Size:        formatFileSize(item.Size()),
	// 				Extension:   Ext(item.Name()),
	// 				Path:        filepath.Join(folder.Path, item.Name()),
	// 				CreatedTime: util.Datetime2DateString(item.ModTime()),
	// 				UpdateTime:  util.Datetime2DateString(item.ModTime()),
	// 			})
	// 		}
	// 	}

	// 	// query content table
	// 	// content, err := model.FindContentByID(folder.ContentID)
	// 	// if err != nil {
	// 	// 	errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
	// 	// 	return
	// 	// }
	// 	// cid := ipfs.ParseCid(content.LocationUrl)
	// 	// ctx, cancel := context.WithCancel(context.Background())
	// 	// defer func() {
	// 	// 	cancel()
	// 	// }()
	// 	// stat, err := ipfs.Stat(ctx, &req.Cfg, cid)
	// 	// if err != nil {
	// 	// 	errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in IPFS: %w", err))
	// 	// 	return
	// 	// }
	// 	// files = append(files, File{
	// 	// 	Name:      folder.Name,
	// 	// 	Type:      "ipfsfile",
	// 	// 	Size:      formatFileSize(stat.Size),
	// 	// 	Extension: "ipfs",
	// 	// 	Path:      content.LocationUrl,

	// 	// 	ContentID:       content.ID,
	// 	// 	ManagedContract: content.ManagedContract,
	// 	// 	KeyID:           content.KeyID,
	// 	// 	LocationUrl:     content.LocationUrl,
	// 	// 	CreatedTime:     util.Datetime2DateString(content.CreatedAt),
	// 	// 	UpdateTime:      util.Datetime2DateString(content.UpdatedAt),
	// 	// })
	// }
	c.JSON(http.StatusOK, ListResponse{
		Folders: folders,
		Files:   files,
	})
}

type CreateRequest struct {
	EncryptType string `json:"encrypt_type"`
	Key         string `json:"key"`
	OriginFile  string `json:"origin_file"`
}

type CreateResponse struct {
	KeyID       int64  `json:"key_id"`
	EncryptFile string `json:"encrypt_file"`
}

func create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}
	input := req.OriginFile
	output := fmt.Sprintf("%s.enc", input)

	if req.EncryptType == "aes" {
		if input == "" || output == "" {
			fmt.Fprintf(os.Stderr, "\033[1;31;40m invalid file path")
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error: invalid file path"))
			return
		}
		if req.Key == "" || len(req.Key) < 16 {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error: invalid aes key"))
			return
		}
		in, err := os.Open(input)
		if err != nil {
			errorResp(c,
				http.StatusInternalServerError,
				xerrors.Errorf("I/O Error: failed to open '%s': %v", input, err))
			return
		}
		out, err := os.Create(output)
		if err != nil {
			errorResp(c,
				http.StatusInternalServerError,
				xerrors.Errorf("I/O Error: failed to create '%s': %v", output, err))
			return
		}
		key, err := encrypt.DeriveKey([]byte(req.Key), in, out)
		if err != nil {
			out.Close()
			os.Remove(out.Name())
			errorResp(c,
				http.StatusInternalServerError,
				xerrors.Errorf("Encrypt err: %v", err))
			return
		}
		record := &model.KeyRecord{
			Password: req.Key,
		}
		keyID, err := record.CreateRecord()
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in DB: %w", err))
			return
		}
		log.Infof("Password saved. [key id is %d ]", keyID)
		cfg := dare.Config{Key: key}
		if _, err := decrypt.AesDecrypt(in, out, cfg); err != nil {
			out.Close()
			os.Remove(out.Name())
			errorResp(c,
				http.StatusInternalServerError,
				xerrors.Errorf("Encrypt err: %v", err))
			return
		}
		log.Infof("Encrypt content finished! [output file is %s]", out.Name())
		c.JSON(http.StatusOK, CreateResponse{
			KeyID:       keyID,
			EncryptFile: out.Name(),
		})
	} else {
		c.JSON(http.StatusOK, CreateResponse{
			KeyID:       -1,
			EncryptFile: "",
		})
	}
}
