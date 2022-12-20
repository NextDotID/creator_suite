package model

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nextdotid/creator_suite/config"
	"github.com/nextdotid/creator_suite/model/contracts"
	"golang.org/x/xerrors"
	"math/big"
)

func CreateAsset(contentId int64, contractAddr string, tokenAddr string, tokenAmount int64) error {
	conn, err := contracts.NewContracts(common.HexToAddress(contractAddr), EthClient)
	if err != nil {
		return xerrors.Errorf("failed to connect the content: %v", err)
	}

	tx_acc := GetTxAccSK()
	transactOps, err := bind.NewKeyedTransactorWithChainID(tx_acc, GetChainID())
	tx, err := conn.CreateAsset(transactOps, uint64(contentId), common.HexToAddress(tokenAddr), big.NewInt(tokenAmount))

	if err != nil {
		return xerrors.Errorf("failed to create the content asset through contract, err:%v, tx:%s", err, tx.Hash().String())
	}
	return nil
}

func IsQualified(addr string, assetId uint64) (bool, error) {
	conn, err := contracts.NewContracts(config.GetSubscriptionContractAddress(), EthClient)
	if err != nil {
		return false, xerrors.Errorf(fmt.Sprintf("failed to connect the content: %v", err))
	}
	return conn.IsQualified(&bind.CallOpts{}, common.HexToAddress(addr), assetId)
}

func GetAssetID(addr string, contentID uint64) (uint64, error) {
	conn, err := contracts.NewContracts(config.GetSubscriptionContractAddress(), EthClient)
	if err != nil {
		return 0, xerrors.Errorf(fmt.Sprintf("failed to connect the content: %v", err))
	}
	return conn.ContentAssetMapping(&bind.CallOpts{}, GetTxAccAddr(), contentID)
}

func GetChainID() *big.Int {
	id, ok := big.NewInt(0).SetString(config.GetChainID(), 10)
	if !ok {
		l.Errorf("fail to parse Chain ID %s", config.GetChainID())
		return big.NewInt(0)
	}
	return id
}

func GetTxAccSK() *ecdsa.PrivateKey {
	skBytes := common.Hex2Bytes(config.GetTxAccConf())
	sk, err := crypto.ToECDSA(skBytes)
	if err != nil {
		l.Errorf("failed to parse paymaster secret key: %v", err)
	}
	return sk
}

func GetTxAccAddr() common.Address {
	skBytes := common.Hex2Bytes(config.GetTxAccConf())
	sk, err := crypto.ToECDSA(skBytes)
	if err != nil {
		l.Errorf("failed to parse paymaster secret key: %v", err)
	}
	return crypto.PubkeyToAddress(sk.PublicKey)
}
