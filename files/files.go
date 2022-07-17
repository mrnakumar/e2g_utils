package Files

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Size(folderPath string) (uint64, error) {
	cmd := exec.Command("du", "-shm", folderPath)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return 0, fmt.Errorf("failed to check size for path '%s'. caused by: '%v'", folderPath, err)
	}

	output := strings.TrimSpace(out.String())
	parts := strings.Split(output, "\t")
	if len(parts) > 0 {
		return strconv.ParseUint(parts[0], 10, 64)
	}
	return 0, fmt.Errorf("failed to check size for path '%s'. unexpected output '%s'", folderPath, output)
}

func ListFiles(suffixes []string, basePath string) ([]fileInfo, error) {
	infos, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}
	var files []fileInfo
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].ModTime().Before(infos[j].ModTime())
	})
	for _, info := range infos {
		if info.Size() > 0 && matchSuffix(suffixes, info.Name()) {
			files = append(files, fileInfo{path: filepath.Join(basePath, info.Name()), size: info.Size(), modTime: info.ModTime()})
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

type fileInfo struct {
	path    string
	size    int64
	modTime time.Time
}
