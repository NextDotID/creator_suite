package model

import (
	"time"

	"golang.org/x/xerrors"
)

const VALID_STATUS = 1
const INVALID_STATUS = 0

type Content struct {
	ID              int64 `gorm:"primarykey"`
	ManagedContract string
	CreatorAddress  string
	KeyID           int64
	LocationUrl     string
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

func CreateRecord(LocateUrl string, managedContract string, keyID int64) (content *Content, err error) {
	c := &Content{}
	c.KeyID = keyID
	c.ManagedContract = managedContract
	c.LocationUrl = LocateUrl
	c.CreatorAddress = GetTxAccAddr().String()
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
