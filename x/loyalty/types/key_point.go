package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PointKeyPrefix is the prefix to retrieve all Point
	PointKeyPrefix = "Point/value/"
)

// PointKey returns the store key to retrieve a Point from the index fields
func PointKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
