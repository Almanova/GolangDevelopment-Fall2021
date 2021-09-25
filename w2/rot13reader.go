package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	n, e := reader.r.Read(b)
	if e == nil {
		for i := range b {
			if b[i] >= 'A' && b[i] <= 'Z' {
				b[i] = table[b[i] - 'A' + 13]
			} else if b[i] >= 'a' && b[i] <= 'z' {
				b[i] = table[(b[i] - 'a' + 39) % 52]
			}
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
