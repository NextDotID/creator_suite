package types

import "math/big"

type Network string

var Networks = struct {
	Ethereum Network
	Goerli   Network
	Polygon  Network
	Mumbai   Network
}{
	Ethereum: "ethereum",
	Goerli:   "goerli",
	Polygon:  "polygon",
	Mumbai:   "mumbai",
}

func (n Network) IsValid() bool {
	switch n {
	case Networks.Ethereum, Networks.Goerli, Networks.Polygon, Networks.Mumbai:
		return true
	}
	return false
}

func (n Network) GetChainID() *big.Int {
	switch n {
	case Networks.Ethereum:
		return big.NewInt(1)
	case Networks.Goerli:
		return big.NewInt(5)
	case Networks.Polygon:
		return big.NewInt(137)
	case Networks.Mumbai:
		return big.NewInt(80001)

	}
	return big.NewInt(0)
}
