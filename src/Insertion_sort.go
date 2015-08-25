package main

import (
	"fmt"
	"time"
)

var begin = []int{3, 1, 5, 6, 7, 2, 4, 9, 8, 0}

func main() {
	timeA := time.Now()
	for i := 0; i < len(begin)-2; i++ {
		for j := i + 1; j < len(begin)-1; j++ {
			if begin[i] < begin[j] {
				begin[i], begin[j] = begin[j], begin[i]
			}
		}
	}
	time.Sleep(time.Second)
	timeB := time.Now()
	fmt.Println(begin)
	fmt.Println(timeB.Sub(timeA.Add(time.Second)))
}
