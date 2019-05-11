package gzip

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
)

func GZipReturn(s string) string {
	var buffer bytes.Buffer
	gZipWriter := gzip.NewWriter(&buffer)

	_, err := gZipWriter.Write([]byte(s))
	if err != nil {
		log.Fatal(err)
	}

	err = gZipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(buffer)
}

func GZipReturnBytes(b []byte) []byte {
	var buffer bytes.Buffer
	gZipWriter := gzip.NewWriter(&buffer)

	_, err := gZipWriter.Write(b)
	if err != nil {
		log.Fatal(err)
	}

	err = gZipWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	return buffer.Bytes()
}