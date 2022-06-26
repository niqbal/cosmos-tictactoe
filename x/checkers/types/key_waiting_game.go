package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WaitingGameKeyPrefix is the prefix to retrieve all WaitingGame
	WaitingGameKeyPrefix = "WaitingGame/value/"
)

// WaitingGameKey returns the store key to retrieve a WaitingGame from the index fields
func WaitingGameKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
