# PhotoArchiver
Photo Archive application written in Go lang.

## Install
go get github.com/rwcarlsen/goexif/exif

## Run
go run main.go -src=./test_pictures -dest=./test_archive



src
    size < max && existing docs in db ?
        for each existing doc
            getStoredHash (existing has hash ? use it : make hash & save)
                hashes match ?
                    insert dupe event (could be many)
        0 match?
            unique
    else
        unique