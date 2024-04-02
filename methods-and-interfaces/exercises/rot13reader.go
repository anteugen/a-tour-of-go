package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(buffer []byte) (n int, err error) {
	n, err = rot.r.Read(buffer)
	for i := 0; i < len(buffer); i++ {
		if (buffer[i] >= 'A' && buffer[i] < 'N') || (buffer[i] >= 'a' && buffer[i] < 'n') {
			buffer[i] += 13
		}  else if (buffer[i] > 'M' && buffer[i] <= 'Z') || (buffer[i] > 'm' && buffer[i] <= 'z') {
			buffer[i] -= 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}