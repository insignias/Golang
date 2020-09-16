package main

import (
	"fmt"
	"io"
	"os"
)

type logReader struct{}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		os.Exit(1)
	}

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	fmt.Println("Error reading the file: ", err)
	// 	os.Exit(1)
	// }
	// fmt.Println(count)
	// fmt.Println(string(data))
	// file.Chmod(0777)
	io.Copy(os.Stdout, file)
	file.Close()

}
