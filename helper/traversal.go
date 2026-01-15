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
	codeStatsData map[string]FolderData,
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
				Traverser(filepath.Join(cwd, entry.Name()), ignoreDirectories, allowedExtensions, codeStatsData)
				// allData.NoOfLines += tempData.NoOfLines
				// allData.NoOfChars += tempData.NoOfChars
			}
		} else {
			fileExtension := filepath.Ext(entry.Name())
			if slices.Contains(allowedExtensions, fileExtension) {
				tempFileData := readFiles(filepath.Join(cwd, entry.Name()))
				fd := codeStatsData[fileExtension]
				fd.NoOfLines += tempFileData.NoOfLines
				fd.NoOfChars += tempFileData.NoOfChars
				codeStatsData[fileExtension] = fd
			}
		}
	}
}
