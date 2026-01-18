// Package tui for better terminal displays
package tui

import (
	"fmt"
	"strings"

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

func DisplayData(codeStatsData map[string]helper.FolderData) {
	rows := make([]Row, 0, len(codeStatsData))
	totalLines := 0

	for ext, data := range codeStatsData {
		rows = append(rows, Row{Ext: ext, Data: data})
		totalLines += data.NoOfLines
	}
	fmt.Printf(
		"%-8s %10s %8s %12s  %s\n",
		"EXT", "LINES", "%", "CHARS", "BAR",
	)
	fmt.Println(strings.Repeat("-", 70))

	for _, row := range rows {
		percent := (float64(row.Data.NoOfLines) / float64(totalLines)) * 100

		fmt.Printf(
			"%-8s %10d %7.1f%% %12d  %s\n",
			row.Ext,
			row.Data.NoOfLines,
			percent,
			row.Data.NoOfChars,
			bar(percent, 25),
		)
	}
}
