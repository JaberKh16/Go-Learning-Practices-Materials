package main

import (
	"fmt"
	"io"
	"os"
)

type ReadError struct {
	Filename string 
	Operation string
}

func (e *ReadError) Error() string {
	return fmt.Sprintf("error reading file %s : %s", e.Filename, e.Operation)
}

func processFile(filename string) error {
	// open
	file, err := os.Open(filename)
	if err != nil {
		return  &os.PathError{ Operation: "open a file", Path: filename, Err: err}
	}
	defer file.Close()

	// read
	buffer := make([]byte, 1024)
	_, err : = file.Read(buffer)
	if err != nil && err != io.EOF {
		return  &ReadError{
			Operation: "read",
			Filename: filename,
		}
	}
	return nil
}


func main(){
	err := processFile("file.txt")
	if err != nil {
		switch e := err.(type) {
			case *os.PathError:
				fmt.Printf("Path error: %s\n", e)
			case *os.ReadError:
				fmt.Printf("Path error: %s\n", e)
			default:
				fmt.Printf("unknown error: %s\n", e)
		}
	
	}
}