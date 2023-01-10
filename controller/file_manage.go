package controller

import (
	"context"
	"fmt"
	"io"
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
	Files       []File `json:"children"`
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

func pathJoin(basePath string, elem ...string) string {
	if strings.HasPrefix(basePath, "./") {
		base := strings.TrimSuffix(basePath, "/")
		return base + "/" + filepath.Join(elem...)
	}
	return filepath.Join(basePath, filepath.Join(elem...))
}

func list(c *gin.Context) {
	req := ListRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	folders := make([]Folder, 0)
	files := make([]File, 0)

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

	// list volumes: STORAGE
	log.Infof("list storage volumes: %s", req.Path)
	list, err := ioutil.ReadDir(req.Path)
	if err != nil {
		log.Infof("I/O error: %v", err)
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
		return
	}
	for _, item := range list {
		if item.IsDir() {
			contentID, err := strconv.ParseInt(item.Name(), 10, 32)
			if err != nil {
				continue
			}
			folder := Folder{
				Name:        item.Name(),
				Type:        "dirs",
				Path:        pathJoin(req.Path, item.Name()),
				ContentID:   contentID,
				CreatedTime: util.Datetime2DateString(item.ModTime()),
				UpdateTime:  util.Datetime2DateString(item.ModTime()),
			}
			children := make([]File, 0)
			f, err := ioutil.ReadDir(folder.Path)
			if err != nil {
				errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
				return
			}
			for _, item := range f {
				if !item.IsDir() {
					children = append(children, File{
						Name:        item.Name(),
						Type:        "localfile",
						Size:        formatFileSize(item.Size()),
						Extension:   Ext(item.Name()),
						Path:        filepath.Join(folder.Path, item.Name()),
						CreatedTime: util.Datetime2DateString(item.ModTime()),
						UpdateTime:  util.Datetime2DateString(item.ModTime()),
					})
				}
			}

			if content, ok := contentMap[contentID]; ok {
				if content.EncryptionType == model.ENCRYPTION_TYPE_AES {
					cid := ipfs.ParseCid(content.LocationUrl)
					ctx, cancel := context.WithCancel(context.Background())
					defer func() {
						cancel()
					}()
					log.Infof("content_id = %d, cid = %s", content.ID, cid)
					stat, err := ipfs.Stat(ctx, &req.Cfg, cid)
					if err != nil {
						errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Error in IPFS: %w", err))
						return
					}
					children = append(children, File{
						Name:      folder.Name,
						Type:      "ipfsfile",
						Size:      formatFileSize(stat.Size),
						Extension: "ipfs",
						Path:      content.LocationUrl,

						ContentID:       content.ID,
						ManagedContract: content.ManagedContract,
						KeyID:           content.KeyID,
						LocationUrl:     content.LocationUrl,
						CreatedTime:     util.Datetime2DateString(content.CreatedAt),
						UpdateTime:      util.Datetime2DateString(content.UpdatedAt),
					})
				}
			}
			folder.Files = children
			folders = append(folders, folder)
		}
	}

	// list host mount path
	list2, err := ioutil.ReadDir(req.Path)
	if err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("I/O error"))
		return
	}
	for _, item := range list2 {
		if !item.IsDir() {
			files = append(files, File{
				Name:        item.Name(),
				Type:        "localfile",
				Size:        formatFileSize(item.Size()),
				Extension:   Ext(item.Name()),
				Path:        pathJoin(req.Path, item.Name()),
				CreatedTime: util.Datetime2DateString(item.ModTime()),
				UpdateTime:  util.Datetime2DateString(item.ModTime()),
			})
		}
	}
	c.JSON(http.StatusOK, ListResponse{
		Folders: folders,
		Files:   files,
	})
}

type CreateRequest struct {
	EncryptType int    `json:"encrypt_type"`
	Key         string `json:"key"`
	OriginFile  string `json:"origin_file"`
	EncryptFile string `json:"encrypt_file"`
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
	output := req.EncryptFile
	if req.EncryptFile == "" {
		output = fmt.Sprintf("%s.enc", input)
	}
	if req.EncryptType == model.ENCRYPTION_TYPE_AES {
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
				xerrors.Errorf("DeriveKey err: %v", err))
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
		if _, err := encrypt.AesEncrypt(in, out, cfg); err != nil {
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

type MoveRequest struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type MoveResponse struct{}

func move(c *gin.Context) {
	req := MoveRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	if req.Src == "" || req.Dst == "" {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error: invalid file path"))
		return
	}

	path := filepath.Dir(req.Dst)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("I/O Error: %v", err))
			return
		}
	}

	log.Infof("Move file src: %s, dst: %s", req.Src, req.Dst)
	err := os.Rename(req.Src, req.Dst)
	if err != nil {
		errorResp(c,
			http.StatusInternalServerError,
			xerrors.Errorf("I/O Error: failed to move %v", err))
		return
	}
	c.JSON(http.StatusOK, MoveResponse{})
}

type CopyRequest struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type CopyResponse struct{}

func copy(c *gin.Context) {
	req := CopyRequest{}
	if err := c.BindJSON(&req); err != nil {
		errorResp(c, http.StatusBadRequest, xerrors.Errorf("Param error"))
		return
	}

	if req.Src == "" || req.Dst == "" {
		errorResp(c, http.StatusInternalServerError, xerrors.Errorf("Param error: invalid file path"))
		return
	}

	path := filepath.Dir(req.Dst)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			errorResp(c, http.StatusInternalServerError, xerrors.Errorf("I/O Error: %v", err))
			return
		}
	}

	log.Infof("Move file src: %s, dst: %s", req.Src, req.Dst)

	fin, err := os.Open(req.Src)
	if err != nil {
		errorResp(c,
			http.StatusInternalServerError,
			xerrors.Errorf("I/O Error: failed to Open %v", err))
		return
	}

	defer fin.Close()

	fout, err := os.Create(req.Dst)
	if err != nil {
		errorResp(c,
			http.StatusInternalServerError,
			xerrors.Errorf("I/O Error: failed to Create %v", err))
		return
	}

	defer fout.Close()

	_, err = io.CopyBuffer(fout, fin, make([]byte, 32*1024))
	if err != nil {
		errorResp(c,
			http.StatusInternalServerError,
			xerrors.Errorf("I/O Error: failed to copy %v", err))
		return
	}
	c.JSON(http.StatusOK, CopyResponse{})
}
