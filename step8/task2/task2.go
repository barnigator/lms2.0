package main

import (
	"context"
	"io"
	"os"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	f, err := os.Open(path)
	if err != nil {
		close(result)
		return
	}
	defer f.Close()

	buffer := make([]byte, 1024)
	var data []byte

	for {
		select {
		case <-ctx.Done():
			close(result)
			return
		default:
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					result <- data
					close(result)
					return
				}
				close(result)
				return
			}
			data = append(data, buffer[:n]...)

		}
	}
}
