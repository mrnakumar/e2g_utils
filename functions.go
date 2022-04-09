package e2g_utils

import (
	"encoding/base64"
	"filippo.io/age"
	"fmt"
	"log"
	"os"
	"strings"
)

type KeyPair struct {
	Public  string
	Private string
}

func Base64Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}

func Base64Decode(input string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func ParsePassword(input string) string {
	str, err := Base64Decode(input)
	if err != nil {
		log.Fatalf("failed to parse password. caused by '%v'", err)
	}
	output := Reverse(str)
	return output
}

func ValidatePath(path *string, flag string) string {
	trimmed := strings.TrimSpace(*path)
	if len(trimmed) == 0 {
		log.Fatalf("flag '%s' is required", flag)
	}
	decoded, err := Base64Decode(trimmed)
	if err != nil {
		log.Fatalf("failed to validate path. error in decoding. error: '%v'", err)
	}
	if _, err := os.Stat(decoded); os.IsNotExist(err) {
		log.Fatalf("invalid '%s'. Path '%s' does not exist.", flag, decoded)
	}
	return decoded
}

func GenerateX25519Identity() KeyPair {
	id, _ := age.GenerateX25519Identity()
	keyPri := id.String()
	keyPub := id.Recipient().String()
	return KeyPair{
		Public:  keyPub,
		Private: keyPri,
	}
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
