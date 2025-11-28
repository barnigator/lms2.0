package main

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, num := range nums {
		ch <- num
	}

	return ch
}
