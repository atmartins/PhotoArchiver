// Package archive copies many assets from one spot to another.
package archive

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atmartins/PhotoArchiver/assets"
	config "github.com/atmartins/PhotoArchiver/config"
	"github.com/atmartins/PhotoArchiver/utils"
)

var constants config.Constants

// All move assets from one spot to another
func All(c config.Constants) {
	constants = c
	fmt.Println("Importing photos from " + constants.DirSrc + " to " + constants.DirArchive)

	// Open the directory
	d, err := os.Open(constants.DirSrc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// When this function (main) is done running, close the directory handler
	defer d.Close()

	// Read the directory contents, recursive-like
	fPathList := []string{}
	err = filepath.Walk(constants.DirSrc, func(path string, f os.FileInfo, err error) error {
		fPathList = append(fPathList, path)
		return nil
	})

	// If no success, give up. Forever. Or until we try again.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	processFiles(fPathList)
}

// Archive moves or copies a single asset
func Archive(src, dst string) (int64, error) {
	return utils.Copy(src, dst)
}

func processFiles(fPathList []string) {
	// Loop through the stuff we found within the directory
	for _, fPath := range fPathList {
		fmt.Println(fmt.Sprintf("rocessing %s", fPath))

		// Open the file since exif.decode (below) expects something with a Reader interface (https://godoc.org/io#Reader)
		dest, skipped, err := processFile(fPath)

		if skipped {
			fmt.Println("\t>> skipping file")
			continue
		}

		if err != nil {
			fmt.Println(fmt.Sprintf("\t!! err: %s", err))
			continue
		}

		fmt.Println(fmt.Sprintf("\t_ moved to %s", dest))
	}
}

func processFile(fPath string) (dest string, skipped bool, err error) {
	var dirDest string

	fInfo, err := os.Stat(fPath)
	if err != nil {
		return "", false, fmt.Errorf("unable to stat fPath, err: %s", err)
	}

	// Make sure the file is regular (https://golang.org/pkg/os/#FileMode)
	if !fInfo.Mode().IsRegular() {
		return "", true, nil
	}

	if assets.IsImage(fPath) {
		dirDest, err = assets.GetImageDirDest(fPath, constants)
		if err != nil {
			// Just move on. Only image files will return exif data.
			// TODO handle assets that lack exif data.
			return "", false, fmt.Errorf("trouble getting file dest %s", err)
		}
		dest = dirDest + fInfo.Name()
	} else if assets.IsVideo(fPath) {
		dirDest = assets.GetVideoDirDest(constants)
		dest = dirDest + fInfo.Name()
	}

	os.MkdirAll(dirDest, 0777)
	_, err = Archive(fPath, dest)

	if err != nil {
		return "", false, fmt.Errorf("unable to Archive. dest: %s err: %s", dest, err)
	}

	return dest, false, nil
}
