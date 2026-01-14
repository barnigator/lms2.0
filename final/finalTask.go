package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

var ValidStatus = map[string]struct{}{
	"В работе":         {},
	"Готово":           {},
	"Не будет сделано": {},
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	resultChan := make(chan error)
	go func(out chan error) {
		data, err := io.ReadAll(r)
		if err != nil {
			out <- err
		}
		strs := strings.Split(string(data), "\n")
		tickets := make([]Ticket, 0, len(strs))

		for _, str := range strs {

			str = strings.TrimSpace(str)
			if str == "" {
				continue
			}

			params := strings.Split(str, "_")

			if len(params) != 4 {
				continue
			}

			if len(params[0]) < 6 {
				continue
			} else if params[0][:6] != "TICKET" {
				continue
			}

			if _, ok := ValidStatus[params[2]]; !ok {
				continue
			}

			timeStamp, err := time.Parse("2006-01-02", params[3])
			if err != nil {
				continue
			}

			tick := Ticket{
				params[0],
				params[1],
				params[2],
				timeStamp,
			}

			if user != nil && tick.User != *user {
				continue
			}

			if status != nil && tick.Status != *status {
				continue
			}

			tickets = append(tickets, tick)

		}
		res, err := json.Marshal(tickets)
		if err != nil {
			out <- err
			return
		}
		_, err = w.Write(res)
		if err != nil {
			out <- err
			return
		}

		out <- nil
	}(resultChan)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(timeout):
		return errors.New("timeout")
	case err := <-resultChan:
		return err
	}
}
