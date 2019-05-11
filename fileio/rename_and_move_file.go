package fileio

import (
	"log"
	"os"
)

func ReNameAndMoveFile() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
