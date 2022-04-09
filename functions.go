package e2g_utils

import (
	"encoding/base64"
	"filippo.io/age"
	"fmt"
)

func Base64Encode(input string) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Printf("Base64 encoding of '%s' is '%s'", input, encoded)
}

func GenerateX25519Identity() {
	id, _ := age.GenerateX25519Identity()
	keyPri := id.String()
	keyPub := id.Recipient().String()
	fmt.Println("------- Generated X25519Identity START --------")
	fmt.Printf("Private Key (bas64 encoded): '%s'\n", base64.StdEncoding.EncodeToString([]byte(keyPri)))
	fmt.Printf("Public Key (bas64 encoded): '%s'\n", base64.StdEncoding.EncodeToString([]byte(keyPub)))
	fmt.Println("------- Generated X25519Identity END --------")
}

func Reverse(str string) string {
	// Get Unicode code points.
	n := 0
	runes := make([]rune, len(str))
	for _, r := range str {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Convert back to UTF-8.
	output := string(runes)
	return output
}
