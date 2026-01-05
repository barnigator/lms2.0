package main

import (
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client := &http.Client{}

		req, err := http.NewRequest("GET", "http://localhost:8081/provideData", nil)
		if err != nil {
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}
		w.Write(body)
	})

	timeoutHandler := http.TimeoutHandler(handler, maxTimeout, "StatusServiceUnavailable")

	http.Handle("/readSource", timeoutHandler)

	http.ListenAndServe(":8080", nil)
}
