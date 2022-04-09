package pkg

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
