package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func collect(ip string, ch chan string, limitChan chan struct{}) {
	defer wg.Done()
	defer close(ch)

	<-limitChan
	ch <- strings.ToUpper(ip)
}

func main() {
	ips := []string{"hello", "world", "china", "go"}
	chs := make([]chan string, len(ips))
	limitChan := make(chan struct{}, len(ips))

	wg.Add(len(ips))
	for i, ip := range ips {
		chs[i] = make(chan string)
		limitChan <- struct{}{}
		go collect(ip, chs[i], limitChan)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	wg.Wait()
}
