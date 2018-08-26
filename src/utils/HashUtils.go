package utils

import (
	"bytes"
	"encoding/binary"
	"log"
	"crypto/sha256"
)

/**
int to hex
 */
func Int2Hex(num int64) []byte {
	buff :=new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

/*
data to hash
 */
func Data2Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
