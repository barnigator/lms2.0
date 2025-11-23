package main

import (
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

func GetTasks(text string, user *string, status *string) []Ticket {
	strs := strings.Split(text, "\n")
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
	return tickets
}
