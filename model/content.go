package model

import (
	"golang.org/x/xerrors"
	"time"
)

type Content struct {
	ID          int64 `gorm:"primarykey"`
	KeyID       int64
	LocationUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Content) TableName() string {
	return "content"
}

func FindContentByID(ID int64) (content *Content, err error) {
	tx := DB.Model(&Content{}).Where("id = ?", ID).Find(&content)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when finding content: %w", err)
	}
	return content, nil
}

func (c *Content) CreateRecord() error {
	tx := DB.Create(c)
	if tx.Error != nil {
		return xerrors.Errorf("error when creating a content record: %w", tx.Error)
	}
	return nil
}
