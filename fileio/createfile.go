package fileio

import (
	"fmt"
	"log"
	"os"
)

var (
	newFile *os.File
	err1    error
)

func CreateFile() {
	newFile, err1 = os.Create("test.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println(newFile)
	if err := newFile.Close(); err != nil {
		fmt.Errorf("The file failed to close: %v", err)
	} else {
		fmt.Println("the file was succesfully created")
	}



}
