package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lebohang0824/aigate/filter"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Missing argument: path to filename")
		os.Exit(1)
	}

	filename := args[1]
	if !fileExists(filename) {
		fmt.Println("Invalid file or file does not exists")
		os.Exit(1)
	}

	content := readFileContent(filename)
	content = filter.Device(content)
	content = filter.Contact(content)
	fmt.Println(content)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}

func readFileContent(filePath string) string {
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return ""
	}

	return string(contentBytes)
}
