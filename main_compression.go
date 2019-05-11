package main

import (
	"fmt"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe/b64"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe/gzip"
	"github.com/Henrikohlsen/Is-105/is105-ica03/pipe/hex"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("No extra arguments provided.")
	}

	file, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}

	x := hex.HexReturnBytes(file)
	b := b64.Base64ReturnBytes(file)
	gz := gzip.GZipReturnBytes(b)
	result := hex.HexReturnBytes(gz)

	i1 := len(file)
	i2 := len(gz)

	LimitPrintln(file, 10, " - Original text")
	LimitPrintln(gz, 10, "- Gzip")
	LimitPrintln(x, 10, " - Hex")
	LimitPrintln(result, 10, "- Result in hex")
	LimitPrintln(b, 10, "- In Base64")

	fmt.Println("-------------------------------------------------------------")

	fmt.Println(len(file), ": The original file length")
	fmt.Println(len(gz), ": Length of the input file after base64 encoding and gzip compression")
	fmt.Println(len(x), ": Original length in hex-encoding")
	fmt.Println(len(result), ": Length of input after pipe in hex-encoding")
	fmt.Println(len(b), ": Length of the input file in base64-encoding")

	fmt.Println("-------------------------------------------------------------")

	fmt.Println("The Compression rate is : ", CompressionRatio(i1, i2), " %")

}

func LimitPrintln(b []byte, limit int, s string) {
	fmt.Println(b[:limit], ": The first", limit, "bytes.", s)
}



func CompressionRatio(old, new int) float64  {
	oldn:= float64(old)
	newn := float64(new)
	percentage := 100 * (1 - (newn / oldn))
	return percentage
	}