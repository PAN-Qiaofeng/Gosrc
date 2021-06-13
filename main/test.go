/*
 * @Description:
 * @version:
 * @Author: pqf
 * @Date: 2021-04-21 15:44:29
 * @LastEditors: pqf
 * @LastEditTime: 2021-06-12 14:11:36
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// func quickSort(tar []int) {
// 	l_tar := len(tar)
// 	for i := 1; i < l_tar; i++ {
// 		if tar[i] < tar[i-1] {
// 			temp := tar[i]
// 			j := i - 1
// 			for ; j >= 0 && temp < tar[j]; j-- {
// 				tar[j+1] = tar[j]
// 			}
// 			tar[j+1] = temp
// 		}
// 	}
// }

func heapSort(nums []int) []int {
	end := len(nums) - 1
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end)
	}
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		sink(nums, 0, end)
	}
	return nums
}

func sink(heap []int, root, end int) {
	for {
		child := root*2 + 1
		if child > end {
			return
		}
		if child < end && heap[child] <= heap[child+1] {
			child++
		}
		if heap[root] > heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}

func quickSort(tar []int, left, right int) {
	if len(tar) > 0 && left < right {
		pivot := partition(tar, left, right)
		quickSort(tar, left, pivot-1)
		quickSort(tar, pivot+1, right)
	}
}
func partition(tar []int, left, right int) int {
	cur := (left + right) / 2
	temp := tar[cur]
	tar[cur], tar[right] = tar[right], tar[cur]
	for j := left; j < len(tar); j++ {
		if tar[j] <= temp {
			tar[j], tar[left] = tar[left], tar[j]
			left++
		}
	}

	return left - 1
}

func insertSort(tar []int) {
	l_tar := len(tar)
	for i := 1; i < l_tar; i++ {
		if tar[i] < tar[i-1] {
			temp := tar[i]
			j := i - 1
			for ; j >= 0 && temp < tar[j]; j-- {
				tar[j+1] = tar[j]
			}
			tar[j+1] = temp
		}
	}

}

var (
	slice = []int{1, 2, 3, 4, 5}
	chE   chan interface{}
)

func main() {
	// file, err := os.OpenFile("practice\\read.txt", os.O_CREATE|os.O_APPEND, 0666)

	// if err != nil {
	// 	return
	// }

	// fileOpen := bufio.NewWriter(file)

	// for i := 0; i < 5; i++ {
	// 	fileOpen.WriteString("我是傻逼\r")
	// }
	test := []int{9, 8, 7, 5, 4, 2, 1}
	heapSort(test)
	// quickSort(test, 0, len(test)-1)
	fmt.Println(test)

	fileRead, err := os.Open("practice\\read.txt")
	if err != nil {
		return
	}
	defer fileRead.Close()
	feToRead := bufio.NewReader(fileRead)
	for {
		str, err := feToRead.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	// fileOpen.Flush()

}
