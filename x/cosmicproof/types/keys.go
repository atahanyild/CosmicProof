package types

const (
	// ModuleName defines the module name
	ModuleName = "cosmicproof"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cosmicproof"
)

var (
	ParamsKey = []byte("p_cosmicproof")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
