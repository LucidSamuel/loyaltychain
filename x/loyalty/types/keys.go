package types

const (
	// ModuleName defines the module name
	ModuleName = "loyalty"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_loyalty"
)

var (
	ParamsKey = []byte("p_loyalty")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
