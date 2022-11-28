package model

import (
	"golang.org/x/xerrors"
	"time"
)

type KeyPair struct {
	ID         int64 `gorm:"primarykey"`
	PublicKey  string
	PrivateKey string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (KeyPair) TableName() string {
	return "key_pair"
}

func FindKeyPairByID(ID int64) (keypair *KeyPair, err error) {
	tx := DB.Model(&KeyPair{}).Where("id = ?", ID).Find(&keypair)
	if tx.Error != nil {
		return nil, xerrors.Errorf("error when finding keypair: %w", err)
	}
	return keypair, nil
}

func (k *KeyPair) CreateRecord() (int64, error) {
	tx := DB.Create(k)
	if tx.Error != nil {
		return 0, xerrors.Errorf("error when creating the keypair: %w", tx.Error)
	}
	return k.ID, nil
}
