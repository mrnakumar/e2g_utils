package e2g_utils

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
			files = append(files, FileInfo{Path: filepath.Join(basePath, info.Name()), Size: info.Size(), ModTime: info.ModTime()})
		}
	}
	return files, nil
}

type FileInfo struct {
	Path    string
	Size    int64
	ModTime time.Time
}
