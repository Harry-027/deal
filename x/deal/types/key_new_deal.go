package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NewDealKeyPrefix is the prefix to retrieve all NewDeal
	NewDealKeyPrefix = "NewDeal/value/"
)

// NewDealKey returns the store key to retrieve a NewDeal from the index fields
func NewDealKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
