package main

import (
	"bufio"
	"errors"
	"os"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	f, err := os.Open(inputFileName)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	res := make([]string, 0)

	scannerDate := bufio.NewScanner(f)

	for scannerDate.Scan() {
		strT := scannerDate.Text()
		t, err := time.Parse("02.01.2006", strT[:10])
		if err != nil {
			return []string{}, err
		}
		if !(t.Before(start) || t.After(end)) {
			res = append(res, strT)
		}
	}

	if len(res) == 0 {
		return res, errors.New("")
	}

	return res, nil

}
