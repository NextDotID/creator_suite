package model

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nextdotid/creator_suite/config"
	"github.com/nextdotid/creator_suite/model/contracts"
	"math/big"
)

func CreateAsset(contentId int64, tokenAddr string, tokenAmount int64) (uint64, error) {
	conn, err := contracts.NewContracts(config.GetSubscriptionContractAddress(), EthClient)
	if err != nil {
		panic(fmt.Sprintf("failed to connect the content: %v", err))
	}

	tx_acc := GetTxAcc()
	transactOps, err := bind.NewKeyedTransactorWithChainID(tx_acc, GetChainID())
	tx, err := conn.CreateAsset(transactOps, uint64(contentId), common.HexToAddress(tokenAddr), big.NewInt(tokenAmount))
	if err != nil {
		panic(fmt.Sprintf("failed to create the content asset through contract: %v", err))
	}

	receipt, err := EthClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil || receipt.Status != 1 {
		panic(fmt.Sprintf("failed to create the content asset through contract: %v", err))
	}

	creator_addr := crypto.PubkeyToAddress(tx_acc.PublicKey)
	return conn.GetAssetId(&bind.CallOpts{}, creator_addr, 3)
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

func GetTxAcc() *ecdsa.PrivateKey {
	skBytes := common.Hex2Bytes("32d438aa3bf89fec159724287e565ebd0dea112afffa132a9772ca0d15fed88b")
	sk, err := crypto.ToECDSA(skBytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse paymaster secret key: %v", err))
	}
	return sk
}
