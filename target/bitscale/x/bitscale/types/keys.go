package types

const (
	// ModuleName defines the module name
	ModuleName = "bitscale"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bitscale"
)

var (
	ParamsKey = []byte("p_bitscale")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
