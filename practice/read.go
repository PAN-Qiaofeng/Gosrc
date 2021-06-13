package main

import (
	"fmt"
)

type Usb struct {
	name string
}

func writeData(intChan chan int) {
	for i := 0; i < 5; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		temp, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(temp)
	}
	exitChan <- true
	close(exitChan)
}
