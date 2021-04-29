package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("practice\\Panqiaofeng.txt")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	fileRead := bufio.NewReader(file)
	for {
		str, err := fileRead.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF {
			fmt.Println("None")
			break
		}
	}
}
