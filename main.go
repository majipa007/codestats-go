package main

import (
	"encoding/json"
	"fmt"
	"os"

	"codestats/helper"
)

func main() {
	// Step 1: Get the current locations
	cwd := helper.GetCwd()

	// Step 2: Get the directories and files to ignore and look for
	data, err := os.ReadFile("codestats.config.json")
	if err != nil {
		panic(err)
	}

	var jsonData map[string][]string

	test := json.Unmarshal(data, &jsonData)
	if test != nil {
		panic(test)
	}

	ignoreDirectories := jsonData["ignore_directories"]
	allowedxtensions := jsonData["allowed_extensions"]

	// Step 3: traverse in the current directory only for the certain directories
	var allFiles []string
	retrivedFiles := helper.Traverser(cwd, ignoreDirectories, allowedxtensions, allFiles)
	for _, data := range retrivedFiles {
		fmt.Println(data)
	}
}
