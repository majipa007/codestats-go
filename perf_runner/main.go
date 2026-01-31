package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"codestats/helper"
)

type runResult struct {
	threads  int
	duration time.Duration
	err      error
}

func main() {
	threadCounts := []int{5000, 1000, 500, 100, 50, 10, 5, 1}

	cwd := helper.GetCwd()

	ignoreDirectories, allowedExtensions, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "config load failed: %v\n", err)
		os.Exit(1)
	}

	results := make([]runResult, 0, len(threadCounts))
	for _, threads := range threadCounts {
		fmt.Printf("\nRunning with %d threads...\n", threads)
		duration, runErr := runWithThreads(
			cwd,
			ignoreDirectories,
			allowedExtensions,
			threads,
		)
		results = append(results, runResult{
			threads:  threads,
			duration: duration,
			err:      runErr,
		})
		if runErr != nil {
			fmt.Fprintf(os.Stderr, "  -> %d threads failed: %v\n", threads, runErr)
		} else {
			fmt.Printf("  -> %d threads finished in %v\n", threads, duration)
		}
	}

	fmt.Println("\nTiming summary:")
	for _, res := range results {
		if res.err != nil {
			fmt.Printf("  %4d threads: failed (%v)\n", res.threads, res.err)
			continue
		}
		fmt.Printf("  %4d threads: %v\n", res.threads, res.duration)
	}
}

func loadConfig() ([]string, []string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, nil, err
	}

	configPath := filepath.Join(home, "codestats", "codestats.config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, nil, err
	}

	var jsonData map[string][]string
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, nil, err
	}

	ignoreDirectories := jsonData["ignore_directories"]
	allowedExtensions := jsonData["allowed_extensions"]
	return ignoreDirectories, allowedExtensions, nil
}

func runWithThreads(
	cwd string,
	ignoreDirectories []string,
	allowedExtensions []string,
	threadCount int,
) (time.Duration, error) {
	startTime := time.Now()

	ch := make(chan helper.FolderData, 100)
	codeStatsData := make(map[string]helper.FolderData)
	var consumerWG sync.WaitGroup
	consumerWG.Add(1)
	go func() {
		defer consumerWG.Done()
		helper.ChannelWriter(ch, codeStatsData)
	}()

	var wg sync.WaitGroup
	sem := make(helper.Semaphore, threadCount)

	helper.Traverser(
		cwd,
		ignoreDirectories,
		allowedExtensions,
		ch,
		&wg,
		sem,
	)
	wg.Wait()
	close(ch)
	consumerWG.Wait()

	return time.Since(startTime), nil
}
