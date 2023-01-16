package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dandascalescu-code/avos-test/files/avos/lzw"
)

func main() {
	curDir, _ := os.Getwd()
	rootDir := filepath.Dir(filepath.Dir(curDir))
	filePath := filepath.Join(rootDir, "data", "01-hello.txt.z")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 12; i++ {
		for j := 7; j >= 0; j-- {
			if data[i]&(1<<uint(j)) != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Print(" ")
	}
	fmt.Println("...")

	newData := lzw.Decompress(data)
	if newData != nil {
		fmt.Println(string(newData))
	}
}
