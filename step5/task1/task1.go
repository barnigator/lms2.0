package main

import (
	"bytes"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			n, err := r.Read(buffer)
			if err != nil {
				if err == io.EOF {
					return false, nil
				}
				return false, err
			}
			if n > 0 {
				if bytes.Contains(buffer[:n], seq) {
					return true, nil
				}
			}
		}
	}
}
