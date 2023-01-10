package types

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
