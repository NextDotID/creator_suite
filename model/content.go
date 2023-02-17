package model

import (
	"time"

	"github.com/nextdotid/creator_suite/types"
	"golang.org/x/xerrors"
)

const VALID_STATUS = 1
const INVALID_STATUS = 0

const ENCRYPTION_TYPE_AES = 1
const ENCRYPTION_TYPE_ECC = 2

type Content struct {
	ID              int64 `gorm:"primarykey"`
	ManagedContract string
	Network         string
	CreatorAddress  string
	EncryptionType  int8 `gorm:"default:1"`
	KeyID           int64
	ContentName     string
	FileExtension   string
	Status          int8 `gorm:"default:1"`
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Content) TableName() string {
	return "content"
}

func FindContentByID(ID int64) (content *Content, err error) {
	tx := DB.Model(&Content{}).Where("id = ? AND status=1", ID).Find(&content)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when finding content: %w", err)
	}
	return content, nil
}

func ListContent() ([]Content, error) {
	contents := make([]Content, 0)
	tx := DB.Model(&Content{}).Where("status = ?", 1).Find(&contents)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when finding content: %w", tx.Error)
	}
	return contents, nil
}

func CreateRecord(managedContract string, keyID int64, encryptionType int8,
	fileExtension string, network types.Network, contentName string, description string) (
	content *Content, err error) {
	c := &Content{}
	c.KeyID = keyID
	c.ContentName = contentName
	c.ManagedContract = managedContract
	c.CreatorAddress = GetTxAccAddr().String()
	c.EncryptionType = encryptionType
	c.FileExtension = fileExtension
	c.Network = string(network)
	c.Description = description
	tx := DB.Create(c)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when creating a content record: %w", tx.Error)
	}
	return c, nil
}

func (c *Content) UpdateLocationUrl(locationUrl string) error {
	tx := DB.Model(&Content{}).Where("id = ?", c.ID).Update("location_url", locationUrl)
	if tx.RowsAffected != 1 || tx.Error != nil {
		return xerrors.Errorf("error when update a content location_url: %w", tx.Error)
	}
	return nil
}

func UpdateAssetID(ID int64, assetID int64) error {
	tx := DB.Model(&Content{}).Where("id = ?", ID).Update("asset_id", assetID)
	if tx.RowsAffected != 1 || tx.Error != nil {
		return xerrors.Errorf("error when update a content asset_id record: %w", tx.Error)
	}
	return nil
}

func (c *Content) UpdateToInvalidStatus(ID int64) error {
	tx := DB.Model(&Content{}).Where("id = ?", ID).Update("status", INVALID_STATUS)
	if tx.RowsAffected != 1 || tx.Error != nil {
		return xerrors.Errorf("error when update a content asset_id record: %w", tx.Error)
	}
	return nil
}
