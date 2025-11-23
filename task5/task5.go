package main

import (
	"io"
	"strings"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	data := make([]byte, 1024)
	_, err := r.Read(data)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(data), string(seq)), nil
}
