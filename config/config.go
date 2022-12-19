package config

import "C"
import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"os"
)

var (
	Viper *viper.Viper
)

// Init initializes config
func Init() {
	Viper = viper.New()
	//Viper.SetConfigName("config_example") // config file name without extension
	Viper.SetConfigType("toml")
	Viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := Viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: cli err:%v \n", err)
		os.Exit(1)
	}
}

type ChainConfig struct {
	ChainID                            string
	RPCServer                          string
	SecretKey                          string
	ContentSubscriptionContractAddress string
}

// GetDatabaseDSN constructs a DSN string for postgresql db driver
func GetDatabaseDSN() string {
	template := "host=%s port=%d user=%s password=%s dbname=%s TimeZone=%s sslmode=disable"
	return fmt.Sprintf(template,
		Viper.GetString("db.host"),
		Viper.GetInt("db.port"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		Viper.GetString("db.tz"),
	)
}
func GetChainID() string {
	return os.Getenv("CHAIN_ID")
}

func GetRPCServer() string {
	return os.Getenv("RPC_SERVER_ON_CHAIN")
}

func GetSubscriptionContractAddress() common.Address {
	return common.HexToAddress(Viper.GetString("chain.subscription_contract_address"))
}

func GetTxAccConf() string {
	return os.Getenv("TX_ACCOUNT")
}
