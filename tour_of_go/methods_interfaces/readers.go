package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Reader_basic() {
	fmt.Println("> Reader basic")
	target := "Hello, world!"
	strReader := strings.NewReader(target)
	bytes := make([]byte, 8)

	for {
		sz, err := strReader.Read(bytes)
		fmt.Printf("n = %v err = %v b = %v\n", sz, err, bytes)
		fmt.Printf("b[:n] = %q\n", bytes[:sz])

		if err == io.EOF {
			break
		}
	}
}

type rot13Reader struct {
	reader io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = (b - 'A') + 'N'
	case 'N' <= b && b <= 'Z':
		b = (b - 'N') + 'A'
	case 'a' <= b && b <= 'm':
		b = (b - 'a') + 'n'
	case 'n' <= b && b <= 'z':
		b = (b - 'n') + 'a'
	}
	return b
}

func (r *rot13Reader) Read(target []byte) (int, error) {
	// transmit data from reader to target
	n, err := r.reader.Read(target)
	
	// transform n element with rot13 cipher
	for i := range(target[:n]){
		target[i] = rot13(target[i])
	}

	// return size, error
	return n, err 
}

func Reader_Wrap() {
	fmt.Println("> Reader wrap")
	// A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
	strReader := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot13Reader := rot13Reader{strReader}
	io.Copy(os.Stdout, &rot13Reader)
	
}
