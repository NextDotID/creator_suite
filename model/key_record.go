package model

import (
	"time"

	"golang.org/x/xerrors"
)

type KeyRecord struct {
	ID        int64 `gorm:"primarykey"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (KeyRecord) TableName() string {
	return "key_record"
}

func FindKeyRecordByID(ID int64) (keyRecord *KeyRecord, err error) {
	tx := DB.Model(&KeyRecord{}).Where("id = ?", ID).Find(&keyRecord)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when finding key_record: %w", err)
	}
	return keyRecord, nil
}

func (k *KeyRecord) CreateRecord() (int64, error) {
	tx := DB.Create(k)
	if tx.Error != nil {
		return 0, xerrors.Errorf("error when creating key_record: %w", tx.Error)
	}
	return k.ID, nil
}
