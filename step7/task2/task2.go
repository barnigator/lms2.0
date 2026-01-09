package main

import (
	"strings"
	"sync"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	cond := sync.NewCond(&sync.Mutex{})
	cond.Signal()
	rightAns := 0
	for i := 0; i < len(answers); i++ {
		select {
		case ans := <-answerCh:
			if strings.EqualFold(ans, answers[i]) {
				rightAns++
			}
		case <-time.After(time.Second):
			<-answerCh
		}
	}

	return rightAns
}
