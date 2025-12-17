package main

import (
	"io/fs"
	"fmt"
	"log"
)

func main() {

}

func loadFile(path string) {
	file, err := fs.Open(path)
	if err != nil {
		return "", err
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return contents, nil
}
