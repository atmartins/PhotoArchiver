package assets

import (
	"fmt"
	"os"

	config "github.com/atmartins/PhotoArchiver/config"
	"github.com/rwcarlsen/goexif/exif"
)

// GetVideoDirDest returns the appropriate place to archive videos to.
func GetVideoDirDest(constants config.Constants) (dest string) {
	return fmt.Sprintf("%s%s/", constants.DirArchive, constants.DirNameVideos)
}

// GetImageDirDest returns the appropriate place to archive images to.
func GetImageDirDest(fPath string, constants config.Constants) (dest string, err error) {
	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		return "", fmt.Errorf("trouble opening file: %s", err)
	}

	// Decode the exif data from the file
	ex, err := exif.Decode(file)
	if err != nil {
		// if no exif data, place in special folder in archive
		return constants.DirArchive + "unknown_date/", nil
	}

	// Determine the dateTaken
	dateTaken, err := ex.DateTime()
	if err != nil {
		return "", fmt.Errorf("trouble determining date taken: %s", err)
	}

	m := int(dateTaken.Month())
	y := dateTaken.Year()

	// Make the directory for year and month, if it doesn't exist already. Grant permissions to the entire planet.
	return fmt.Sprintf("%s%s/%d/%d/", constants.DirArchive, constants.DirNamePhotos, y, m), nil
}
