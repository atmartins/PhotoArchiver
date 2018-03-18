package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

type Constants struct {
	DirSrc        string `json:"DirSrc"`
	DirArchive    string `json:"DirArchive"`
	DirNamePhotos string `json:"DirNamePhotos"`
	DirNameVideos string `json:"DirNameVideos"`
}

var constants Constants

func main() {
	constants = loadConstants()
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

func loadConstants() Constants {
	raw, err := ioutil.ReadFile("./constants.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Constants
	json.Unmarshal(raw, &c)
	return c
}

func processFiles(fPathList []string) {
	// Loop through the stuff we found within the directory
	for _, fPath := range fPathList {
		fmt.Println(fmt.Sprintf("Processing %s", fPath))

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
		return "", false, fmt.Errorf("Unable to stat fPath, err: %s", err)
	}

	// Make sure the file is regular (https://golang.org/pkg/os/#FileMode)
	if !fInfo.Mode().IsRegular() {
		return "", true, nil
	}

	if isImage(fPath) {
		file, err := os.Open(fPath)
		defer file.Close()
		if err != nil {
			return "", false, fmt.Errorf("Trouble opening file: %s", err)
		}
		dirDest, err = getImageDirDest(file)
		if err != nil {
			// Just move on. Only image files will return exif data.
			// TODO handle images that lack exif data.
			return "", false, fmt.Errorf("Trouble getting file dest %s", err)
		}
		dest = dirDest + fInfo.Name()
	} else if isVideo(fPath) {
		dirDest = getVideoDirDest()
		dest = dirDest + fInfo.Name()
	}

	os.MkdirAll(dirDest, 0777)
	_, err = archive(fPath, dest)

	if err != nil {
		return "", false, fmt.Errorf("Unable to archive. dest: %s err: %s", dest, err)
	}

	return dest, false, nil
}

func getVideoDirDest() (dest string) {
	return fmt.Sprintf("%s%s/", constants.DirArchive, constants.DirNameVideos)
}

func getImageDirDest(f io.Reader) (dest string, err error) {
	// Decode the exif data from the file
	ex, err := exif.Decode(f)
	if err != nil {
		// if no exif data, place in special folder in archive
		return constants.DirArchive + "unknown_date/", nil
	}

	// Determine the dateTaken
	dateTaken, err := ex.DateTime()
	if err != nil {
		return "", fmt.Errorf("Trouble determining date taken: %s", err)
	}

	m := int(dateTaken.Month())
	y := dateTaken.Year()

	// Make the directory for year and month, if it doesn't exist already. Grant permissions to the entire planet.
	return fmt.Sprintf("%s%s/%d/%d/", constants.DirArchive, constants.DirNamePhotos, y, m), nil
}

func archive(src, dst string) (int64, error) {
	return Copy(src, dst)
}

func isImage(fPath string) bool {
	switch strings.ToLower(filepath.Ext(fPath)) {
	case
		".jpeg",
		".jpg",
		".png",
		".tiff",
		".cr2",
		".dng",
		".gif",
		".bmp":
		return true
	}
	return false
}

func isRawImage(fPath string) bool {
	switch strings.ToLower(filepath.Ext(fPath)) {
	case
		".tiff",
		".cr2",
		".dng":
		return true
	}
	return false
}

func isVideo(fPath string) bool {
	switch strings.ToLower(filepath.Ext(fPath)) {
	case
		".3gp",
		".mts",
		".mkv",
		".mp4",
		".mov":
		return true
	}
	return false
}

/*Copy file from src to dst
 * Courtesy of user edap on http://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
 */
func Copy(src, dst string) (int64, error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	srcFileStat, err := srcFile.Stat()
	if err != nil {
		return 0, err
	}

	if !srcFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}
