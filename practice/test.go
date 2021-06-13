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

func main_01() {
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

	sliceByte := []rune(testString)

	sliceByte[0] = '我'

	fmt.Println(testString)

	slByString := string(sliceByte)

	fmt.Printf("%v \n", slByString)

	newString := testString + slByString

	fmt.Println(newString)

	fmt.Printf("%T \n", newString)

	var twoDeArray [2][3]int

	fmt.Println(twoDeArray)

	for i := 0; i < len(twoDeArray); i++ {
		for j := 0; j < len(twoDeArray[i]); j++ {
			fmt.Println(twoDeArray[i][j])
		}
	}

	tranInt := fmt.Sprintf("%d", a[0])

	fmt.Printf("%T \n", tranInt)

	mapTest := make(map[string]string)

	mapTest2 := map[string]string{
		"面试1": "面试2",
		"腾讯1": "腾讯2",
	}

	fmt.Println(mapTest["面试1"], mapTest["腾讯2"])

	delete(mapTest, "腾讯2")

	fmt.Println(mapTest["腾讯2"], mapTest["面试1"])
	fmt.Println(mapTest2["面试1"], mapTest2["腾讯1"])

	panic(e)

}
