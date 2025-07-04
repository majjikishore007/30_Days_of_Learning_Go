package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func pickRandom(input []string) string {
	if len(input) == 0 {
		return ""
	}
	key := rand.Intn(len(input))
	return input[key]
}

func main() {
	// read all the files from the source
	fortunePath := "/usr/share/fortune"

	var files []string

	filepath.Walk(fortunePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".dat" {
			return nil
		}
		files = append(files, path)
		return nil
	})

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No fortune files found\n")
		os.Exit(1)
	}

	// randomly pick one fie
	randmonFile := pickRandom(files)

	// Read the file and make a slice of quotes
	file, err := os.ReadFile(randmonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}
	quotes := strings.Split(string(file), "%")
	// randomly pick a quote
	quote := strings.TrimSpace(pickRandom(quotes))
	if quote != "" {
		fmt.Println(quote)
	}
}
