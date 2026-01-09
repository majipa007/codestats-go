package main

import (
	"fmt"

	"codestats/helper"
)

func main() {
	// Step 1: Get the current locations
	cwd := helper.GetCwd()

	// Step 2: traverse in the current directory
	// testArray := helper.Traverser(cwd)
	helper.Traverser(cwd)
	fmt.Println(cwd)
}
