package assets

import (
	"crypto/md5"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// IsRawImage returns true if fPath ends in camera raw img file extension.
func IsRawImage(fPath string) bool {
	switch strings.ToLower(filepath.Ext(fPath)) {
	case
		".tiff",
		".cr2",
		".dng":
		return true
	}
	return false
}

// IsImage returns true if fPath ends in any img file extension.
func IsImage(fPath string) bool {
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

// IsVideo returns true if fPath ends in any video file extension.
func IsVideo(fPath string) bool {
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

// ComputeHash returns the md5 hash of a file
func ComputeHash(fPath string) (hashStr string, err error) {
	f, err := os.Open(fPath)

	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	hashStr = string(h.Sum(nil))
	return
}
