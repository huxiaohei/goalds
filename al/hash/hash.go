package hash

import (
	"bytes"
	"crypto/sha512"
	"encoding/binary"
)

func GetHashInts(seed []byte, n int) []uint64 {
	date := seed
	var hashInts []uint64
	for len(hashInts) < n {
		data := Hash512(date)
		tmp := make([]uint64, len(data)/8)
		buf := bytes.NewBuffer(data)
		binary.Read(buf, binary.BigEndian, &tmp)
		hashInts = append(hashInts, tmp...)
	}
	return hashInts[:n]
}

func Hash512(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}
