package main

import (
	"fmt"

	"github.com/crossphoton/drive-encrypt/src"
)

var pass = "password"

func main_() {

	files, err := src.ListAllFiles(pass)
	if err != nil {
		panic(err)
	}

	fmt.Println("djkjkabdk", files)

	err = src.NewFile(make([]byte, 56), src.File{
		Name: "test",
		Type: "dummy",
		Size: 56,
	}, pass)
	if err != nil {
		panic(err)
	}

	files, err = src.ListAllFiles(pass)
	if err != nil {
		panic(err)
	}

	fmt.Println(files)
}
