package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		panic(err.Error())
		return "", err
	}
	return fmt.Sprintf("%x",uuid[:]),nil
}
