package fileio

import (
	"log"
	"os"
)

func DeleteFile() {
	err := os.Remove("test2.txt")
	if err != nil {
		log.Fatal(err)
	}


}
