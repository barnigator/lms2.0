package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	ctxT, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req.WithContext(ctxT))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiR := &APIResponse{string(body), resp.StatusCode}

	return apiR, nil
}
