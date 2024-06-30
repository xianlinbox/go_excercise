package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader	struct{}

func (fileReader) readFromFile(filename string)  {
	file,err := os.Open(filename)

	if err != nil {
		fmt.Println("Error:", err)	
		os.Exit(1)
	}	

	io.Copy(os.Stdout, file)
}