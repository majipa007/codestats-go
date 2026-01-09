package helper

import (
	"fmt"
	"os"
)

type fileData struct {
	fileType  string
	noOfLines int32
}

func Traverser(cwd string) {
	folderContent, err := os.ReadDir(cwd)
	if err != nil {
		panic(err)
	}
	// var array []fileData
	for _, entry := range folderContent {
		fmt.Println(entry)
		if entry.IsDir() {
			Traverser(cwd + "/" + entry.Name())
		}
		// x := fileData{"test", 5}
		// array = append(array, x)
	}
	// return array
}
