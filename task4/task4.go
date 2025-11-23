package main

import "io"

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, 1024)

	readByte, err := r.Read(data)
	if err != nil {
		return err
	}

	data = data[:min(readByte, int(n))]

	_, err = w.Write(data)
	return err
}
