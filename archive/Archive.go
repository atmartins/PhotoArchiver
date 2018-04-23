// Package archive copies many assets from one spot to another.
package archive

import (
	"fmt"

	config "github.com/atmartins/PhotoArchiver/config"
	"github.com/atmartins/PhotoArchiver/files"
)

var constants config.Constants

// All works through all files in src, attempting to archive to dest
func All(src string, dest string) (stats *ArchStats, err error) {
	fmt.Println(fmt.Sprintf("archiving from %s to %s", src, dest))

	fPathList, err := files.GetFilePaths(src)
	if err != nil {
		return nil, err
	}

	stats, err = processFiles(fPathList)
	if err != nil {
		return nil, err
	}

	return
}

func processFiles(fPathList []string) (stats *ArchStats, err error) {
	stats = &ArchStats{len(fPathList), 0, 0, 0, 0}

	// Loop through the stuff we found within the directory
	for _, fPath := range fPathList {
		isDupe, isUnique, err := singleFile(fPath)

		if err != nil {
			stats.NumError++
			fmt.Println(fmt.Sprintf("\t_ error %s", fPath))
			continue
		}

		if isDupe {
			stats.NumDupes++
			fmt.Println(fmt.Sprintf("\t_ dupe %s", fPath))
		}

		if isUnique {
			stats.NumUnique++
			fmt.Println(fmt.Sprintf("\t_ unique %s", fPath))
		}

		stats.NumSuccess++
		fmt.Println(fmt.Sprintf("\t_ processed %s", fPath))
	}

	return
}
