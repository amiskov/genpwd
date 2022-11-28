package cryptosource

import (
	"encoding/binary"
	"log"

	crand "crypto/rand"
)

// Combining `math/rand` with a source that gets its data from `crypto/rand`.
// https://yourbasic.org/golang/crypto-rand-int/

type CryptoSource struct{}

func New() *CryptoSource {
	return &CryptoSource{}
}

func (s CryptoSource) Seed(seed int64) {}

func (s CryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s CryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return
}
