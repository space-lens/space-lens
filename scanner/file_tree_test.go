package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileTreeScanner(t *testing.T) {

	fileTree := NewFilesTree()

	// split every line
	fmt.Println("Loading data")
	start := time.Now()

	err := filepath.Walk("/Users/alfiankan/perkuliahan/skripsi/devskripsi/web-new/idresearch-indexs", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		cols := strings.Split(path, "/")
		filePathPart := []string{"/"}
		filePathPart = append(filePathPart, cols...)

		fileTree.AddFile(filePathPart, int(info.Size()))

		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded", time.Now().Sub(start))

	paths := fileTree.GetPathList("/Users/alfiankan/perkuliahan/skripsi/devskripsi/web-new/idresearch-indexs")

	for _, f := range paths {
		fmt.Println(f.Name, f.Size)
	}

	assert.True(t, len(paths) > 0)
}
