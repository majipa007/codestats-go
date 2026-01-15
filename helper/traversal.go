package helper

import (
	"os"
	"path/filepath"
	"slices"
)

type FolderData struct {
	NoOfLines int
	NoOfChars int
}

func Traverser(cwd string,
	ignoreDirectories []string,
	allowedExtensions []string,
	allData *FolderData,
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
				Traverser(filepath.Join(cwd, entry.Name()), ignoreDirectories, allowedExtensions, allData)
				// allData.NoOfLines += tempData.NoOfLines
				// allData.NoOfChars += tempData.NoOfChars
			}
		} else {
			fileExtention := filepath.Ext(entry.Name())
			if slices.Contains(allowedExtensions, fileExtention) {
				tempFileData := readFiles(filepath.Join(cwd, entry.Name()))

				allData.NoOfLines += tempFileData.NoOfLines
				allData.NoOfChars += tempFileData.NoOfChars
			}
		}
	}
}
