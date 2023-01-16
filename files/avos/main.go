package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	curDir, _ := os.Getwd()
	rootDir := filepath.Dir(filepath.Dir(curDir))
	filePath := filepath.Join(rootDir, "data", "01-hello.txt.z")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T", data)
	}
}
