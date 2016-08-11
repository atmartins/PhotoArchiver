package main

import (
    "fmt"
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
            f, err := os.Open(dirname + a.Name())
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
            s := fmt.Sprintf("%s was taken in %d/%d",a.Name(), m, y)
            fmt.Println(s)
        }
    }
}
