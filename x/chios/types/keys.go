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
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("chios-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
