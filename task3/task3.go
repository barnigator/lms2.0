package main

import "strings"

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(p []byte) (n int, err error) {
	str := strings.ToUpper(string(p))
	u.UpperString = str
	return len(p), err
}
