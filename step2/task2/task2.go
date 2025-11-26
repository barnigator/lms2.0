package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)

	for i := 0; fileScanner.Scan(); i++ {
		if i == lineNum {
			return fileScanner.Text()
		}
	}
	return ""
}
