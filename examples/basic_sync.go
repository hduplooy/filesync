package main

import (
	filesync "github.com/hduplooy/filesync"
	"log"
)

func main() {
	err := filesync.SyncFolders(&filesync.Config{
		SourcePath:      "D:/tmp/static",
		DestinationPath: "D:/tmp/static2",
	})

	if err != nil {
		log.Fatal(err)
	}
}
