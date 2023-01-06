package scanner

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Represent Each Tree Node
type FileNode struct {
	Name     string
	FullPath string
	Size     int
	Childs   map[string]*FileNode
}

type FileTree struct {
	RootTree *FileNode
}

func NewFilesTree() *FileTree {
	return &FileTree{
		RootTree: &FileNode{
			Name:   "/",
			Childs: make(map[string]*FileNode),
		},
	}
}

// Add path file to tree
func (tree *FileTree) AddFile(fullPath []string, size int) {
	filePath := fullPath[1:]
	currentTree := tree.RootTree
	for {
		if len(filePath) == 0 {
			break
		}
		if currentTree.Childs[filePath[0]] == nil {
			currentTree.Childs[filePath[0]] = &FileNode{
				Name:     filePath[0],
				FullPath: strings.Join(fullPath[1:], "/"),
				Size:     size,
				Childs:   make(map[string]*FileNode),
			}
		}
		// part path already exist
		// pas next to next node
		currentTree = currentTree.Childs[filePath[0]]
		filePath = filePath[1:]
	}
}

// Private: getFolderSize full depth
func (tree *FileTree) getFolderSize(path string) (folderSize int64) {
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				folderSize += info.Size()
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return
}

// Get Files or Folders listing
func (tree *FileTree) GetPathList(path string) map[string]*FileNode {
	if path == "/" {
		return tree.RootTree.Childs
	}

	pathPart := strings.Split(path, "/")
	currNode := tree.RootTree
	for _, f := range pathPart {
		currNode = currNode.Childs[f]
	}

	for _, f := range currNode.Childs {
		if len(f.Childs) > 0 {
			// is folder
			f.Size = int(tree.getFolderSize(f.FullPath))
		}
	}
	return currNode.Childs
}
