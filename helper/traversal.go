package helper

import (
	"os"
	"path/filepath"
	"slices"
)

// type fileData struct {
// 	fileType  string
// 	noOfLines int32
// }

func Traverser(cwd string,
	ignoreDirectories []string,
	allowedExtensions []string,
	allfiles []string,
) []string {
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
				tempArray := Traverser(cwd+"/"+entry.Name(), ignoreDirectories, allowedExtensions, allfiles)
				allfiles = append(allfiles, tempArray...)
			}
		} else {
			fileExtention := filepath.Ext(entry.Name())
			if slices.Contains(allowedExtensions, fileExtention) {
				allfiles = append(allfiles, entry.Name())
			}
		}
	}
	return allfiles
}
