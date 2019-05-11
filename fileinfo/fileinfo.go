package fileinfo

import (
	"fmt"
	"log"
	"os"
)

var (
	fileinfo os.FileInfo
	err      error
)

//this is a function to read files, it takes any kind of string and reads it.
func FileReader(name string) {
	//runs lstat to get info about file..
	fileinfo, err = os.Lstat(name)
	if err != nil {
		log.Fatal(err)
	}

	fileMode :=fileinfo.Mode()
	appOnly := fileMode&os.ModeAppend != 0
	var device bool
	var unixChar  bool
	var unixBlock bool
	//checks whether file is Device
	if fileMode&os.ModeDevice != 0 {
		device = true
		//checks whether file is CharDevice
		if fileMode&os.ModeCharDevice != 0 {
			unixChar = true
			//if not CharDevice & if not Device then its block
		} else {
			unixBlock = true
		}
	}
	symbolicLink := fileMode&os.ModeSymlink != 0

	fmt.Println("file name:", fileinfo.Name())
	fmt.Println("Size in bytes", fileinfo.Size())
	bytes := fileinfo.Size()
	nByte := float64(bytes)
	//gjør om til float før casting
	kibibytes := nByte / 1024 //casting to float64
	mibibytes := kibibytes / 1024
	gibibytes := mibibytes / 1024
	fmt.Println("KibiBytes:", kibibytes)
	fmt.Println("MibiBytes:", mibibytes)
	fmt.Println("GibiBytes:", gibibytes)
	fmt.Println("Permissions:", fileinfo.Mode())
	fmt.Printf("Permissions codes: %#o\n", fileinfo.Mode().Perm())
	fmt.Println("Last modified:", fileinfo.ModTime())
	fmt.Println("Is a directory:", fileinfo.IsDir())
	fmt.Println("Is a regular file:", fileinfo.Mode().IsRegular())
	fmt.Println("Is append only:", appOnly)
	fmt.Println("Is a device:",device)
	fmt.Println("Is a Unix character device:", unixChar)
	fmt.Println("Is a Unix block device:", unixBlock)
	fmt.Println("Is a symbolic link:", symbolicLink)

	//hvis blokk device, da er det ikke char device.
}

