package main

import "time"

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	timer := time.AfterFunc(100*time.Millisecond, func() {
		close(stop)
	})
	defer timer.Stop()

	isPrime := true

	for i := 2; i <= N; i++ {
		select {
		case <-stop:
			close(prime_nums)
			return
		default:
			for y := 2; y <= i/2; y++ {
				if i%y == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				select {
				case <-stop:
					close(prime_nums)
					return
				default:
					prime_nums <- i
				}
			}
			isPrime = true
		}
	}
	close(prime_nums)
}
