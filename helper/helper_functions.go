// Package helper is for getting functions that we might need for the main funciotn
package helper

import (
	"bufio"
	"os"
)

func GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

func readFiles(path string) FolderData {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close() //nolint:errcheck

	scanner := bufio.NewScanner(f)
	lines := 0
	chars := 0

	for scanner.Scan() {
		lines++
		chars += len(scanner.Text())
	}

	return FolderData{lines, chars}
}
