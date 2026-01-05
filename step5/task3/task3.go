package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	URL        string
	Data       string
	StatusCode int
	Err        error
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	out := make(chan *APIResponse)
	apiRs := make([]*APIResponse, 0, len(urls))

	for _, url := range urls {
		go func(url string, in chan<- *APIResponse) {
			ctxT, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			client := &http.Client{}

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				in <- &APIResponse{url, "", 0, err}
				return
			}
			resp, err := client.Do(req.WithContext(ctxT))
			if err != nil {
				in <- &APIResponse{url, "", 0, err}
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				in <- &APIResponse{url, "", 0, err}
				resp.Body.Close()
				return
			}
			defer resp.Body.Close()

			apiR := &APIResponse{url, string(body), resp.StatusCode, nil}

			in <- apiR
		}(url, out)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case apiR := <-out:
			apiRs = append(apiRs, apiR)
		case <-ctx.Done():
			for j := len(apiRs); j < len(urls); j++ {
				apiR := &APIResponse{
					URL: urls[j],
					Err: ctx.Err(),
				}
				apiRs = append(apiRs, apiR)
			}
		}
	}

	return apiRs
}
