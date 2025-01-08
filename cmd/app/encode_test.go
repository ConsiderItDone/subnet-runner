package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// stringToByte32 converts a hex string to a [32]byte array
func stringToByte32(s string) ([32]byte, error) {
	var b [32]byte
	bytes, err := hex.DecodeString(s[2:]) // Remove the "0x" prefix and decode the hex string
	if err != nil {
		return b, err
	}
	copy(b[:], bytes)
	return b, nil
}

// byte32ToString converts a [32]byte array to a hex string
func byte32ToString(b [32]byte) string {
	return "0x" + hex.EncodeToString(b[:])
}

func TestStringToByte32(t *testing.T) {
	hexStr := "0x311a06aafd1a4e714e223a05e90ba6fcb85006493cf880272abf5db51f68161c"

	byteArray, err := stringToByte32(hexStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	decodedString := byte32ToString(byteArray)
	fmt.Println(decodedString)
	if hexStr != decodedString {
		t.Errorf("Expected %s, got %s", hexStr, decodedString)
	}

	remoteBlockchainID := [32]byte{49, 26, 6, 170, 253, 26, 78, 113, 78, 34, 58, 5, 233, 11, 166, 252, 184, 80, 6, 73, 60, 248, 128, 39, 42, 191, 93, 181, 31, 104, 22, 28}

	str := byte32ToString(remoteBlockchainID)
	fmt.Println(str)
}
