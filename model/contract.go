package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nextdotid/creator_suite/config"
	"github.com/nextdotid/creator_suite/model/contracts"
	"golang.org/x/xerrors"
)

//func CreateAsset(contentId int64, contractAddr string, tokenAddr string, tokenAmount *big.Int, network types.Network) error {
//	conn, err := contracts.NewContracts(common.HexToAddress(contractAddr), EthClient)
//	if err != nil {
//		return xerrors.Errorf("failed to connect the content: %v", err)
//	}
//	tx_acc := GetTxAccSK()
//	transactOps, err := bind.NewKeyedTransactorWithChainID(tx_acc, network.GetChainID())
//	tx, err := conn.CreateAsset(transactOps, uint64(contentId), common.HexToAddress(tokenAddr), tokenAmount)
//	if err != nil {
//		return xerrors.Errorf("failed to create the content asset through contract, err:%v, tx:%v", err, tx)
//	}
//	return nil
//}

func IsQualified(contractAddr string, addr string, assetId uint64) (bool, error) {
	conn, err := contracts.NewContracts(common.HexToAddress(contractAddr), EthClient)
	if err != nil {
		return false, xerrors.Errorf(fmt.Sprintf("failed to connect the content: %v", err))
	}
	return conn.IsQualified(&bind.CallOpts{}, common.HexToAddress(addr), assetId)
}

func GetAssetID(contractAddr string, addr string, contentID uint64) (uint64, error) {
	conn, err := contracts.NewContracts(common.HexToAddress(contractAddr), EthClient)
	if err != nil {
		return 0, xerrors.Errorf(fmt.Sprintf("failed to connect the content: %v", err))
	}
	return conn.ContentAssetMapping(&bind.CallOpts{}, common.HexToAddress(addr), contentID)
}

//func GetTxAccSK() *ecdsa.PrivateKey {
//	skBytes := common.Hex2Bytes(config.GetTxAccConf())
//	sk, err := crypto.ToECDSA(skBytes)
//	if err != nil {
//		l.Errorf("failed to parse paymaster secret key: %v", err)
//	}
//	return sk
//}

func GetTxAccAddr() common.Address {
	skBytes := common.Hex2Bytes(config.GetTxAccConf())
	sk, err := crypto.ToECDSA(skBytes)
	if err != nil {
		l.Errorf("failed to parse paymaster secret key: %v", err)
	}
	return crypto.PubkeyToAddress(sk.PublicKey)
}
