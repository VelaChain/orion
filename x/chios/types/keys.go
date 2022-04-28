package types

import (
	"sort"
	"fmt"

	//sdk "github.com/cosmos/cosmos-sdk/types"
)

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

func PrefixKey(key string) []byte {
	return []byte(key)
} 

// func GetPoolKeyFromAssets(pa PoolAssets) []byte {
// 	// sort assets by name
// 	sort.Sort(pa)
// 	// create name w/ scheme: a-b-...-n
// 	name := []byte(pa.Assets[0].Symbol)
// 	for i, a := range pa.Assets {
// 		if i > 0 {
// 			name = append(name, []byte(fmt.Sprintf("-%s", a.Symbol))...)
// 		}
// 	}
// 	return append(KeyPoolPrefix, name...)
// }

func GetPoolNameFromAssets(pa PoolAssets) string {
	// sort assets by name
	sort.Sort(pa)
	// create name w/ scheme: a-b-...-n
	name := pa.Assets[0].Symbol
	for i, a := range pa.Assets {
		if i > 0 {
			name += fmt.Sprintf("-%s", a.Symbol)
		}
	}
	return name
}

func GetPoolKeyFromPoolName(poolName string) []byte {
	poolKeyByte := []byte(fmt.Sprintf("-%s", poolName))
	return append(KeyPoolPrefix, poolKeyByte...)
}

// key for specific provider in a given pool
// use this key for adding a liquidity provider
func GetProviderKey(poolName string, creator string) []byte {
	// return []byte of pool prefix, symbol, creator
	provKeyByte := []byte(fmt.Sprintf("/%s/%s", poolName, creator))
	return append( KeyProviderPrefix, provKeyByte... )
}

// key for provider all providers in a given pool
// use this key to get providers for pool from providers store
func GetProvidersKey(poolName string) [] byte {
	provKeyByte := []byte(fmt.Sprintf("/%s", poolName))
	return append(KeyProviderPrefix, provKeyByte...)
}

// key for all providers in all pools
func GetAllProvidersKey() []byte {
	return KeyProviderPrefix
}
