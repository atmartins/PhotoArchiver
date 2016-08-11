package main

import (
    "fmt"
    "os"
    "path/filepath"
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
            // Print the file's name and size
            fmt.Println(a.Name(), a.Size(), "bytes")
        }
    }
}
