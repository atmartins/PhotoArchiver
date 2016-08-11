package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "github.com/rwcarlsen/goexif/exif"
)

func main() {
    // Define the directory of my pictures
    dirname := "." + string(filepath.Separator) + "test_pictures" + string(filepath.Separator)

    // Open the directory
    d, err := os.Open(dirname)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    // When this function (main) is done running, close the directory handler
    defer d.Close()

    // Read the directory contents
    fi, err := d.Readdir(-1)

    // If no success, give up. Forever. Or until we try again.
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Loop through the stuff we found within the directory
    for _, a := range fi {
        // Make sure the file is regular (https://golang.org/pkg/os/#FileMode)
        if a.Mode().IsRegular() {
            // Open the file since exif.decode (below) expects something with a Reader interface (https://godoc.org/io#Reader)
            src := dirname + a.Name()
            f, err := os.Open(src)

            defer f.Close()

            if err != nil {
                // If we couldn't open it, let's just move on. Not sure if we're getting things like '.' and '..' as contents of our directory
                continue
            }

            // Decode the exif data from the file
            ex, err := exif.Decode(f)
            if err != nil {
                // Just move on. Only image files will return exif data.
                // TODO handle images that lack exif data.
                continue
            }

            // Determine the dateTaken
            dateTaken, err := ex.DateTime()
            if err != nil {
                continue
            }

            m := int(dateTaken.Month())
            y := dateTaken.Year()

            // Make the directory for year and month, if it doesn't exist already. Grant permissions to the entire planet.
            destDir := fmt.Sprintf("archive/%d/%d", y, m)
            os.MkdirAll(destDir, 0777)

            // File destination.
            dest := fmt.Sprintf("archive/%d/%d/%s", y, m, a.Name())

            // Attempt to copy the file to the archive.
            _, err = Copy(src, dest)
            if err != nil {
                fmt.Println("Unable to copy " + src + " to " + dest)
                continue
            }

            fmt.Println("Copied " + src + " to " + dest)
        }
    }
}

// Courtesy of user edap on http://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
func Copy(src, dst string) (int64, error) {
    src_file, err := os.Open(src)
    if err != nil {
        return 0, err
    }
    defer src_file.Close()

    src_file_stat, err := src_file.Stat()
    if err != nil {
        return 0, err
    }

    if !src_file_stat.Mode().IsRegular() {
        return 0, fmt.Errorf("%s is not a regular file", src)
    }

    dst_file, err := os.Create(dst)
    if err != nil {
        return 0, err
    }
    defer dst_file.Close()
    return io.Copy(dst_file, src_file)
}
