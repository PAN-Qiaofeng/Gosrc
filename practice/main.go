package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func judgeCir(test []LinkNode) bool {
	if len(test) == 0 {
		return true
	}
	slow := test[0].Next
	fast := test[0].Next.Next
	for slow != fast {
		if fast == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	//计算环
	calcuLen := &test[0]
	length := 0
	for calcuLen != slow {
		calcuLen = calcuLen.Next
		slow = slow.Next
		length++
	}
	fmt.Println(length)

	return true
}

// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

// 示例 1:

// 输入: coins = [1, 2, 5], amount = 11
// 输出: 3
// 解释: 11 = 5 + 5 + 1
// 示例 2:

// 输入: coins = [2], amount = 3
// 输出: -1
// 说明:
// 你可以认为每种硬币的数量是无限的。

func calcuSum(coin []int, target int) {
	counts := make([]int, len(coin))
	sum := 0
	for sum != target {
		for i := 0; i < len(counts); i++ {
			if target < sum+coin[i] {
				continue
			} else if target > sum+coin[i] {
				coin[i]++
			} else {
				coin[i]++
				break
			}
		}

	}

}

func main() {

	nodeSlice := make([]LinkNode, 10)

	for i := 0; i < 10; i++ {
		nodeSlice[i] = LinkNode{
			i,
			nil,
		}
	}

	for i := 0; i < 10; i++ {
		if i != 9 {
			nodeSlice[i].Next = &nodeSlice[i+1]
		} else {
			nodeSlice[i].Next = &nodeSlice[3]
		}

	}

	flag := judgeCir(nodeSlice)
	if flag {
		fmt.Println("有环")
	} else {
		fmt.Println("无环")
	}

}
