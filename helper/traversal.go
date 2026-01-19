package helper

import (
	"os"
	"path/filepath"
	"slices"
	"sync"
)

type FolderData struct {
	FileExtension string
	NoOfLines     int
	NoOfChars     int
}

func Traverser(cwd string,
	ignoreDirectories []string,
	allowedExtensions []string,
	ch chan FolderData,
	wg *sync.WaitGroup,
) {
	// getting the folder content in an array
	folderContent, err := os.ReadDir(cwd)
	if err != nil {
		panic(err)
	}

	// looping each foldderContent
	for _, entry := range folderContent {
		// checking if the folder ocntent is a directory or a file
		if entry.IsDir() {
			// trying to ignore thing inside the ignoreDirectories
			if !slices.Contains(ignoreDirectories, string(entry.Name())) {
				Traverser(filepath.Join(cwd, entry.Name()),
					ignoreDirectories,
					allowedExtensions,
					ch,
					wg,
				)
			}
		} else {
			fileExtension := filepath.Ext(entry.Name())
			if slices.Contains(allowedExtensions, fileExtension) {
				wg.Add(1)
				go readFiles(filepath.Join(cwd, entry.Name()), fileExtension, ch, wg)
			}
		}
	}
}
