package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("file.txt")
	reader := io.Reader(file)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		fmt.Printf("Read n=%v, err={%v}, buffer={%v}, content={%v}\n", n, err, buffer, string(buffer[:n]))
		if err == io.EOF {
			break
		}
	}
	file.Close()
}
