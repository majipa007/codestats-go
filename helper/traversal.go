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

type Semaphore chan struct{}

func (s Semaphore) Acquire() { s <- struct{}{} }
func (s Semaphore) Release() { <-s }

func Traverser(cwd string,
	ignoreDirectories []string,
	allowedExtensions []string,
	ch chan FolderData,
	wg *sync.WaitGroup,
	sem Semaphore,
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
				wg.Add(1)
				go func(path string) {
					defer wg.Done()
					sem.Acquire()
					defer sem.Release()
					Traverser(
						path,
						ignoreDirectories,
						allowedExtensions,
						ch,
						wg,
						sem,
					)
				}(filepath.Join(cwd, entry.Name()))
			}
		} else {
			fileExtension := filepath.Ext(entry.Name())
			if slices.Contains(allowedExtensions, fileExtension) {
				wg.Add(1)
				go func(path, ext string) {
					defer wg.Done()
					sem.Acquire()
					defer sem.Release()
					readFiles(path, ext, ch)
				}(filepath.Join(cwd, entry.Name()), fileExtension)
			}
		}
	}
}
