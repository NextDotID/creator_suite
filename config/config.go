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

	Viper.SetConfigType("toml")
	//viper.AddConfigPath(".")
	Viper.AddConfigPath("./config/") // config file path
	//viper.AutomaticEnv()             // read value ENV variable

	err := Viper.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file err:%v \n", err)
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
		Viper.GetString("db.user"),
		Viper.GetString("db.password"),
		Viper.GetString("db.db_name"),
		Viper.GetString("db.tz"),
	)
}
func GetChainID() string {
	return Viper.GetString("chain.id")
}

func GetRPCServer() string {
	return Viper.GetString("chain.rpc_server")
}

func GetSubscriptionContractAddress() common.Address {
	return common.HexToAddress(Viper.GetString("chain.subscription_contract_address"))
}
