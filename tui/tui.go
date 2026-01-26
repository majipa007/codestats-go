// Package tui handles terminal output formatting.
package tui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"codestats/helper"
)

type Row struct {
	Ext  string
	Data helper.FolderData
}

func bar(percent float64, width int) string {
	blocks := int((percent / 100) * float64(width))
	return strings.Repeat("█", blocks)
}

func DisplayData(codeStatsData map[string]helper.FolderData, time_taken time.Duration) {
	rows := make([]Row, 0, len(codeStatsData))

	var totalLines int
	var totalChars int

	for ext, data := range codeStatsData {
		rows = append(rows, Row{
			Ext:  ext,
			Data: data,
		})
		totalLines += data.NoOfLines
		totalChars += data.NoOfChars
	}

	// Sort by descending line count
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Data.NoOfLines > rows[j].Data.NoOfLines
	})

	// Header
	fmt.Println()
	fmt.Println("Code Statistics Summary")
	fmt.Println(strings.Repeat("─", 73))
	fmt.Printf(
		"%-8s %10s %8s %12s  %s\n",
		"EXT", "LINES", "%", "CHARS", "DISTRIBUTION",
	)
	fmt.Println(strings.Repeat("-", 73))

	// Rows
	for _, row := range rows {
		percent := (float64(row.Data.NoOfLines) / float64(totalLines)) * 100

		fmt.Printf(
			"%-8s %10d %7.1f%% %12d  %s\n",
			row.Ext,
			row.Data.NoOfLines,
			percent,
			row.Data.NoOfChars,
			bar(percent, 30),
		)
	}

	// Footer (totals)
	fmt.Println(strings.Repeat("-", 73))
	fmt.Printf(
		"%-8s %10d %7.1f%% %12d\n",
		"TOTAL",
		totalLines,
		100.0,
		totalChars,
	)
	fmt.Println(strings.Repeat("─", 73))
	fmt.Printf(
		"%-8s %10.3f ms (%.3f s)",
		"TIME",
		time_taken.Seconds()*1000,
		time_taken.Seconds(),
	)
	fmt.Println()
}
