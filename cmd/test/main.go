package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

/*
select s_name
from
() a
left join
() b
on

*/
func main() {
	fmt.Println(Fibonacci(4))
}

func HideFirstCharacter(s string) string {
	s1 := []rune(s)
	s1[0] = '*'
	return string(s1)
}

func ParseString() string {
	sli := strings.Split("2020-05-16 19:20:34|user.login|name=Charles&location=Beijing&device=iPhone", "|")
	result := strings.Split(sli[2], "&")

	rt := make(map[string]string)
	for _, v := range result {
		key := strings.Split(v, "=")[0]
		value := strings.Split(v, "=")[1]
		rt[key] = value
	}

	bytes, _ := json.Marshal(rt)
	return string(bytes)
}

func GetMinAbs(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	if arr[0] >= 0 {
		return arr[0]
	} else if arr[len(arr)-1] <= 0 {
		return arr[len(arr)-1]
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] < 0 {
			low = mid + 1
		} else if arr[mid] > 0 {
			if arr[mid]*arr[mid-1] < 0 {
				return int(math.Min(math.Abs(float64(arr[mid-1])), float64(arr[mid])))
			}
			high = mid - 1
		} else {
			return arr[mid]
		}
	}
	return -1
}

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
