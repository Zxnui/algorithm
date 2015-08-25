package main

import (
	"algorithm/common"
	"fmt"
	"strings"
)

func main() {
	help()

	stack := [10]string{}
	var stackId int
	for {
		in := common.ReadDateFromInput()

		input := strings.Split(in, " ")
		switch input[0] {
		case "push":
			if stackId == 10 {
				fmt.Println("栈上溢")
			} else {
				stack[stackId] = input[1]
				stackId++
			}
		case "pop":
			if stackId == 0 {
				fmt.Println(stack)
			} else {
				stackId--
				fmt.Println(stack[stackId])
				stack[stackId] = ""
			}
		case "show":
			fmt.Println(stack)
		case "quit":
			goto C
		default:
			help()

		}
	}
C:
}

func help() {
	fmt.Println("--栈操作\n命令：push ..	入栈\n命令：pop	出栈\n命令：show	展示栈内元素\n命令：quit	退出")
}
