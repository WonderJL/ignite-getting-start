package types

const (
	// ModuleName defines the module name
	ModuleName = "ignitegettingstart"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ignitegettingstart"
)

var (
	ParamsKey = []byte("p_ignitegettingstart")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
