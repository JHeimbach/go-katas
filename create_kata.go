package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("no name given")
	}

	var name = strings.Join(os.Args[1:], "-")

	folderPath := filepath.Join(".", name)
	err := os.Mkdir(folderPath, os.ModePerm)

	if err != nil {
		log.Fatalf("could not create dir %q", folderPath)
	}

	createGoFile("kata.go", folderPath)
	createGoFile("kata_test.go", folderPath)
}

func createGoFile(name string, folderPath string) {
	filename := filepath.Join(folderPath, name)
	err := ioutil.WriteFile(filename, []byte("package kata \n\n"), os.ModePerm)
	if err != nil {
		log.Fatalf("could not create file %q", filename)
	}
}
