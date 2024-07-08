package types

const (
	// ModuleName defines the module name
	ModuleName = "ytubclambda"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ytubclambda"
)

var (
	ParamsKey = []byte("p_ytubclambda")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
