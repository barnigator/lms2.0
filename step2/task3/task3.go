package main

import "os"

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	f, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	offset := startpos
	_, err = f.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	buffer := make([]byte, 1024)

	n, err := f.Read(buffer)
	if err != nil {
		return err
	}

	err = os.WriteFile(outFileName, buffer[:n], 0666)
	if err != nil {
		return err
	}

	return nil
}
