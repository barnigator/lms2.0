package main

import (
	"io"
	"os"
)

func ModifyFile(filename string, pos int, val string) {
	file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	defer file.Close()

	file.Seek(int64(pos), io.SeekStart)

	file.WriteString(val)

}
