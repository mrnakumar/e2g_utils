package e2g_utils

import (
	"bytes"
	"encoding/base64"
	"filippo.io/age"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type KeyPair struct {
	Public  string
	Private string
}

func Decrypt(filePath string, privateKey *age.X25519Identity) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	decrypted, err := age.Decrypt(bytes.NewBuffer(content), privateKey)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(decrypted)
	return buf.Bytes(), err
}

func Base64DecodeWithKill(input string, errorName string) string {
	decoded, err := Base64Decode(input)
	if err != nil {
		log.Fatalf("failed to decode '%s'. caused by '%v'", getFirstNonEmpty(errorName, input), err)
	}
	return decoded
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

func ParsePassword(input string, errorName string) string {
	str, err := Base64Decode(input)
	if err != nil {
		log.Fatalf("failed to parse '%s'. caused by '%v'", getFirstNonEmpty(errorName, input), err)
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

func getFirstNonEmpty(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	return a
}

// list files

type FileInfo struct {
	path    string
	size    int64
	modTime time.Time
}

func ListFiles(suffixes []string, basePath string) ([]FileInfo, error) {
	infos, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}
	var files []FileInfo
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].ModTime().Before(infos[j].ModTime())
	})
	for _, info := range infos {
		if info.Size() > 0 && matchSuffix(suffixes, info.Name()) {
			files = append(files, FileInfo{path: filepath.Join(basePath, info.Name()), size: info.Size(), modTime: info.ModTime()})
		}
	}
	return files, nil
}

func matchSuffix(suffixes []string, fileName string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(fileName, suffix) {
			return true
		}
	}
	return false
}
