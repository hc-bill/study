package main

import "time"

func main() {
	arr1 := []int{1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}

	ch1 := make(chan struct{})
	ch2 := make(chan struct{}, 1)

	ch2 <- struct{}{}
	go func() {
		for i := 0; i < len(arr1); i++ {
			<-ch2
			println(arr1[i])
			ch1 <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < len(arr2); i++ {
			<-ch1
			println(arr2[i])
			ch2 <- struct{}{}
		}
	}()

	time.Sleep(1 * time.Second)
}
