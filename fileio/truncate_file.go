package fileio
//oppgave 3a
import (
	"log"
	"os"
)

func TruncateFile() {
	err := os.Truncate("test.txt", 100)
	if err != nil {
		log.Fatal(err)
	}


}
