package model

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nextdotid/creator_suite/config"
	"github.com/nextdotid/creator_suite/model/contracts"
	"math/big"
)

func CreateAsset(contentId int64, contractAddr string, tokenAddr string, tokenAmount int64) (uint64, error) {
	conn, err := contracts.NewContracts(common.HexToAddress(contractAddr), EthClient)
	if err != nil {
		panic(fmt.Sprintf("failed to connect the content: %v", err))
	}

	tx_acc := GetTxAccSK()
	transactOps, err := bind.NewKeyedTransactorWithChainID(tx_acc, GetChainID())
	_, err = conn.CreateAsset(transactOps, uint64(contentId), common.HexToAddress(tokenAddr), big.NewInt(tokenAmount))
	if err != nil {
		panic(fmt.Sprintf("failed to create the content asset through contract: %v", err))
	}

	creator_addr := crypto.PubkeyToAddress(tx_acc.PublicKey)
	return conn.GetAssetId(&bind.CallOpts{}, creator_addr, uint64(contentId))
}

func IsQualified(addr string, assetId uint64) (bool, error) {
	conn, err := contracts.NewContracts(config.GetSubscriptionContractAddress(), EthClient)
	if err != nil {
		panic(fmt.Sprintf("failed to connect the content: %v", err))
	}
	return conn.IsQualified(&bind.CallOpts{}, common.HexToAddress(addr), assetId)
}

func GetChainID() *big.Int {
	id, ok := big.NewInt(0).SetString(config.GetChainID(), 10)
	if !ok {
		panic(fmt.Sprintf("failed to parse chain id: %v", config.GetChainID()))
	}
	return id
}

func GetTxAccSK() *ecdsa.PrivateKey {
	skBytes := common.Hex2Bytes(config.GetTxAccConf())
	sk, err := crypto.ToECDSA(skBytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse paymaster secret key: %v", err))
	}
	return sk
}
