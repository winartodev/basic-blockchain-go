package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntegerToHex(input int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, input)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
