package encryption

import (
	"bytes"
	"filippo.io/age"
	"fmt"
	"github.com/mrnakumar/e2g_utils"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Decoder struct {
	identity *age.X25519Identity
}

func CreateDecoder(privateKeyFilePath string) (Decoder, error) {
	privateKey, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		return Decoder{}, fmt.Errorf("failed to read file '%s'. Caused by : '%v'", privateKeyFilePath, err)
	}

	trimmed := strings.TrimSuffix(string(privateKey), "\n")
	decoded, err := e2g_utils.Base64Decode(trimmed)
	if err != nil {
		return Decoder{}, err
	}
	identity, err := age.ParseX25519Identity(decoded)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	return Decoder{identity: identity}, err
}

func (e Decoder) Decrypt(data string) ([]byte, error) {
	r, err := age.Decrypt(strings.NewReader(data), e.identity)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data")
	}
	out := &bytes.Buffer{}
	if _, err := io.Copy(out, r); err != nil {
		return nil, fmt.Errorf("failed to decrypt")
	}
	return out.Bytes(), nil
}

func CreateEncryptor(recipientKeyPath string) (Encryptor, error) {
	recipientKey, err := ioutil.ReadFile(recipientKeyPath)
	if err != nil {
		return Encryptor{}, fmt.Errorf("failed to read file '%s'. Caused by : '%v'", recipientKeyPath, err)
	}
	trimmed := strings.TrimSuffix(string(recipientKey), "\n")
	decoded, err := e2g_utils.Base64Decode(trimmed)
	if err != nil {
		return Encryptor{}, fmt.Errorf("failed to decode recepient key path. Caused by: '%v'", err)
	}
	publicKey, err := age.ParseX25519Recipient(decoded)
	return Encryptor{recipient: publicKey}, err
}

func (e Encryptor) Encrypt(data []byte) ([]byte, error) {
	out := &bytes.Buffer{}
	w, err := age.Encrypt(out, e.recipient)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	err = w.Close()
	return out.Bytes(), err
}

type Encryptor struct {
	recipient *age.X25519Recipient
}
