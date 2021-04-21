/*
 * @Description:
 * @version:
 * @Author: pqf
 * @Date: 2021-04-21 15:44:29
 * @LastEditors: pqf
 * @LastEditTime: 2021-04-21 21:50:50
 */
package main

import (
	"errors"
	"fmt"
)

var slice = []int{1, 2, 3, 4, 5}

func main() {
	var a [5]int
	a[0] = 1
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v \r", b)

	slice := a[1:3]
	fmt.Printf("%v", slice)

	var sliceB []int = make([]int, 5)

	var sliceC []int = []int{1, 2, 3, 4, 5}

	e := errors.New("test error")

	sliceB = append(sliceB, 1, 2, 3, 4)

	fmt.Println(sliceB, sliceC)

	sliceTest := make([]int, 5)

	copy(sliceTest, sliceB)

	fmt.Printf("%v \r", sliceTest)

	testString := "We are in the same world"

	fmt.Println(testString)

	panic(e)

}
