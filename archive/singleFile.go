package archive

import "os"

// singleFile processes a single file
func singleFile(fPath string) (isDupe bool, isUnique bool, err error) {
	// src
	// 	existing docs in db same size ?
	//      size > max ?
	//          dupe
	//      else
	//    		for each existing doc
	// 			getStoredHash (existing has hash ? use it : make hash & save)
	// 				hashes match ?
	// 					insert dupe event (could be many)
	//              else
	//                  unique
	// 	else
	// 		unique
	fi, err := os.Open(fPath)
	if err != nil {
		return false, false, err
	}
	defer fi.Close()

	stats, err := fi.Stat()
	if err != nil {
		return false, false, err
	}
	size := stats.Size()
	if size < constants.MaxFileSizeBytesToTest {
		isUnique = true
		// processUnique(fi)
	}

	return
	//
	//
	// TODO
	// open/stat file, get size in bytes
	//

	// var dirDest string

	// fInfo, err := os.Stat(fPath)
	// if err != nil {
	// 	return "", false, fmt.Errorf("unable to stat fPath, err: %s", err)
	// }

	// // Make sure the file is regular (https://golang.org/pkg/os/#FileMode)
	// if !fInfo.Mode().IsRegular() {
	// 	return "", true, nil
	// }

	// if assets.IsImage(fPath) {
	// 	dirDest, err = assets.GetImageDirDest(fPath, constants)
	// 	if err != nil {
	// 		// Just move on. Only image files will return exif data.
	// 		// TODO handle assets that lack exif data.
	// 		return "", false, fmt.Errorf("trouble getting file dest %s", err)
	// 	}
	// 	dest = dirDest + fInfo.Name()
	// } else if assets.IsVideo(fPath) {
	// 	dirDest = assets.GetVideoDirDest(constants)
	// 	dest = dirDest + fInfo.Name()
	// }

	// os.MkdirAll(dirDest, 0777)
	// _, err = Archive(fPath, dest)

	// if err != nil {
	// 	return "", false, fmt.Errorf("unable to Archive. dest: %s err: %s", dest, err)
	// }

	// return dest, false, nil
}
