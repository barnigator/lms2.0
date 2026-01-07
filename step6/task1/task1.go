package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetMark(name string) (int, error) {
	client := &http.Client{}

	url := fmt.Sprintf("http://localhost:8082/mark?name=%s", name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return 0, errors.New("student not found")
	}

	if resp.StatusCode == 500 {
		return 0, errors.New("server error")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		return 0, err
	}

	num, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return num, nil
}

func Compare(name1, name2 string) (string, error) {
	mark1, err := GetMark(name1)
	if err != nil {
		return "", err
	}
	mark2, err := GetMark(name2)
	if err != nil {
		return "", err
	}

	switch {
	case mark1 > mark2:
		return ">", nil
	case mark1 < mark2:
		return "<", nil
	default:
		return "=", nil

	}
}
