// Package helper is for getting functions that we might need for the main funciotn
package helper

import (
	"os"
)

func GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}
