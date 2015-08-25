package main

import (
	"algorithm/common"
	"fmt"
	"strings"
)

type List struct {
	Prev string
	Next string
	Key  string
}

var Head string
var Length int

func main() {
	help()
	ListAll := make([]List, 100)

	for {
		in := common.ReadDateFromInput()
		input := strings.Split(in, " ")
		switch input[0] {
		case "push":
			if Length == 0 {
				ListAll[0].Key = input[1]
				ListAll[0].Next = "nil"
				ListAll[0].Prev = "nil"
				Length++
			} else {
				ListAll[Length].Key = input[1]
				ListAll[Length].Prev = ListAll[Length-1].Key
				ListAll[Length].Next = "nil"
				ListAll[Length-1].Next = ListAll[Length].Key
				Length++
			}
		case "delete":
			if Length == 0 {
				fmt.Println("链表为空")
			} else {
				for id, one := range ListAll {
					if one.Key == input[1] {
						if one.Next == "nil" && one.Prev == "nil" {
							Length--
							ListAll[0] = List{"", "", ""}
						} else if one.Prev == "nil" {
							for i := 0; i < Length; i++ {
								ListAll[i] = ListAll[i+1]
							}
							Length--
							ListAll[Length] = List{"", "", ""}
							ListAll[Length-1].Prev = "nil"
						} else if one.Next == "nil" {
							Length--
							ListAll[Length] = List{"", "", ""}
							ListAll[Length-1].Next = "nil"
						} else {
							for i := id; i <= Length; i++ {
								ListAll[id] = ListAll[id+1]
							}
							Length--
						}
						break
					}
				}
			}
		case "selete":
			var bo bool
			if Length == 0 {
				fmt.Println("链表为空")
			}
			for _, one := range ListAll {
				if one.Key == input[1] {
					fmt.Println(one)
					bo = true
				}
			}
			if !bo {
				fmt.Println("无此数据")
			}
		case "show":
			if Length == 0 {
				fmt.Println("链表为空")
			} else {
				fmt.Println()
				ListAllShow := ""
				for i := 0; i < Length; i++ {
					ListAllShow = ListAllShow + "prev:" + ListAll[i].Prev + ",key:" + ListAll[i].Key + ",next:" + ListAll[i].Next + "\n"
				}
				ListAllShow = strings.TrimSpace(ListAllShow)
				fmt.Println(ListAllShow + "\n")
			}
		case "quit":
			goto C
		default:
			help()
		}
	}
C:
}

func help() {
	fmt.Println("--链表操作\n命令：push ..	链表插入\n命令：delete ..	链表删除\n命令：selete ..	搜素链表\n命令：show	展示所有链表元素\n命令：quit	退出")
}
