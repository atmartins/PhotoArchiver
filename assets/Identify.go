package assets

import (
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
