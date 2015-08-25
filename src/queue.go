package main

import (
	"algorithm/common"
	"fmt"
	"strings"
)

func main() {
	help()

	queue := [10]string{}
	var queueId int
	for {
		in := common.ReadDateFromInput()

		input := strings.Split(in, " ")
		switch input[0] {
		case "push":
			if queueId == 10 {
				fmt.Println("队上溢")
			} else {
				queue[queueId] = input[1]
				queueId++
			}
		case "pop":
			if queueId == 0 {
				fmt.Println("队下溢")
			} else {
				queueId--
				fmt.Println(queue[0])
				for i := 0; i < queueId; i++ {
					queue[i] = queue[i+1]
				}
				queue[queueId] = ""
			}
		case "show":
			fmt.Println(queue)
		case "quit":
			goto C
		default:
			help()

		}
	}
C:
}

func help() {
	fmt.Println("--队列操作\n命令：push ..	入队\n命令：pop	出队\n命令：show	展示队内元素\n命令：quit	退出")
}
