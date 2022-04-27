package types

const (
	// ModuleName defines the module name
	ModuleName = "chios"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_chios"

	// Version defines the current version the IBC module supports
	Version = "chios-1"

	// PortID is the default port id that module binds to
	PortID = "chios"

	// default fee for swaps and exiting
	PoolDefaultFee = "0.003"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = []byte{0x00}
	// key for storing the pools map
	KeyPoolPrefix = []byte{0x01}	
	// key for storing lps
	KeyProviderPrefix = []byte{0x02}
	// key for pool count
	KeyPoolCount = []byte{0x03}
)

func GetPoolKeyFromAssets(assets PoolAssets) []byte {
	// sort assets by name
	sort.Sort(assets)
	// create name w/ scheme: a-b-...-n
	name := []byte(assets[0].Symbol)
	for i, a := range assets {
		if i > 0 {
			name += []byte(fmt.Sprintf("-%s", a.Symbol))
		}
	}
	return append(KeyPoolPrefix, name)
}

func GetPoolKeyFromPoolName(poolName string) []byte {
	return append(KeyPoolPrefix,[]byte(poolName))
}

// key for specific provider in a given pool
// use this key for adding a liquidity provider
func GetProviderKey(poolName string, creator string) []byte {
	// return []byte of pool prefix, symbol, creator
	return append( KeyProviderPrefix, []byte( append( poolName, creator ) ) )
}

// key for provider all providers in a given pool
// use this key to get providers for pool from providers store
func GetProvidersKey(poolName string) [] byte {
	return append(KeyProviderPrefix, []byte(poolName) )
}


