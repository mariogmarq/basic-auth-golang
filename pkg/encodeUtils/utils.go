package encodeUtils

import "crypto/sha256"

func Encode(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	b := make([]byte, hasher.Size())
	hasher.Sum(b)
	return b
}
