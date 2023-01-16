package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dandascalescu-code/avos-test/files/avos/lzw"
)

const OUTPUT_DIR = "output"

func main() {
	curDir, _ := os.Getwd()
	rootDir := filepath.Dir(filepath.Dir(curDir))

	err := os.MkdirAll(filepath.Join(rootDir, OUTPUT_DIR), os.ModePerm)
	if err != nil {
        fmt.Println(err)
        return
    }

	files := []string {"01-hello.txt.z", "02-book.txt.z", "03-lyrics.txt.z", "04-icon.png.z"}

	for _, fileName := range files {
		fmt.Printf("\n%v\n", fileName)

		filePath := filepath.Join(rootDir, "data", fileName)
		ext := filepath.Ext(fileName)
		outPath := filepath.Join(rootDir, OUTPUT_DIR, strings.TrimSuffix(fileName, ext))

		// Read input
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			continue
		}
	
		// Decompress data
		output := lzw.Decompress(data)
		if output == "" {
			fmt.Println("Empty output.")
			continue
		}

		// Write output
		err = os.WriteFile(outPath, []byte(output), 0644)
        if err != nil {
            fmt.Println(err)
            continue
        }

		fmt.Printf("Output to %v\n", strings.TrimSuffix(fileName, ext))
	}
}
