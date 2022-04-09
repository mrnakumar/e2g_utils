package main

import (
	"flag"
	"mrnakumar.com/e2g-utils/pkg"
)

type utilFlags struct {
	generateKey bool
	toEncode    string
}

func main() {
	utilFlags := parseFlags()
	if utilFlags.generateKey {
		pkg.GenerateX25519Identity()
	}
	if len(utilFlags.toEncode) > 0 {
		pkg.Base64Encode(utilFlags.toEncode)
	}
}

func parseFlags() utilFlags {
	generateKey := flag.String("generate_X25519_key", "", "generate key? [true | false]")
	toEncode := flag.String("to_encode", "", "input string to decode.")
	flag.Parse()
	shouldGenerateKey := false
	if *generateKey == "true" {
		shouldGenerateKey = true
	}
	return utilFlags{
		toEncode:    *toEncode,
		generateKey: shouldGenerateKey,
	}
}
