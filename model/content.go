package model

import (
	"github.com/nextdotid/creator_suite/types"
	"golang.org/x/xerrors"
	"time"
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
	LocationUrl     string
	FileExtension   string
	Status          int8 `gorm:"default:1"`
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

func CreateRecord(locateUrl string, managedContract string, keyID int64, encryptionType int8, fileExtension string, network types.Network) (content *Content, err error) {
	c := &Content{}
	c.KeyID = keyID
	c.ManagedContract = managedContract
	c.LocationUrl = locateUrl
	c.CreatorAddress = GetTxAccAddr().String()
	c.EncryptionType = encryptionType
	c.FileExtension = fileExtension
	c.Network = string(network)
	tx := DB.Create(c)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when creating a content record: %w", tx.Error)
	}
	return c, nil
}

func (c *Content) UpdateAssetID(ID int64, assetID int64) error {
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
