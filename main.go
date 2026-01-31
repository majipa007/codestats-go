package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"codestats/helper"
	"codestats/tui"
)

func main() {
	startTime := time.Now()
	// Step 1: Get the current locations
	cwd := helper.GetCwd()
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Step 2: Get the directories and files to ignore and look for
	configPath := filepath.Join(home, "codestats", "codestats.config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	var jsonData map[string][]string
	codeStatsData := make(map[string]helper.FolderData)
	test := json.Unmarshal(data, &jsonData)
	if test != nil {
		panic(test)
	}

	ignoreDirectories := jsonData["ignore_directories"]
	allowedxtensions := jsonData["allowed_extensions"]

	ch := make(chan helper.FolderData, 100)
	var consumerWG sync.WaitGroup
	consumerWG.Add(1)
	go func() {
		defer consumerWG.Done()
		helper.ChannelWriter(ch, codeStatsData)
	}()
	// Step 3: traverse in the current directory only for the certain directories and get the allowed_extensions
	// Step 4: get the array of the particular dtype
	var wg sync.WaitGroup
	sem := make(helper.Semaphore, 500)
	helper.Traverser(cwd, ignoreDirectories, allowedxtensions, ch, &wg, sem)

	// wait for workers
	wg.Wait()

	// now safe to close the channel
	close(ch)
	consumerWG.Wait()
	tui.DisplayData(codeStatsData, time.Since(startTime))
}
