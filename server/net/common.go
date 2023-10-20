package net

import (
	"crypto/rand"
	"fmt"
)

func generateUUID() (uuid string) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])

	return
}
