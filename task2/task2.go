package main

import "io"

func ReadString(r io.Reader) (string, error) {
	data := make([]byte, 1024)
	bytesRead, err := r.Read(data)

	if err == io.EOF {
		return "", nil
	}

	return string(data[:bytesRead]), err
}
