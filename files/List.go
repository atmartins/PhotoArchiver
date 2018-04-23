package files

import (
	"os"
	"path/filepath"
)

// GetFilePaths returns list of strings of files found in src
func GetFilePaths(src string) (fPathList []string, err error) {
	// Open the directory
	d, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	// When this function (main) is done running, close the directory handler
	defer d.Close()

	// Read the directory contents, recursive-like
	err = filepath.Walk(src, func(path string, f os.FileInfo, err error) error {
		fPathList = append(fPathList, path)
		return nil
	})

	// If no success, give up. Forever. Or until we try again.
	if err != nil {
		return nil, err
	}

	return
}
