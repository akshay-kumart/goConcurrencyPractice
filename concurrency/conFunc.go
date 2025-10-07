package concurrency

import (
	"fmt"
	"time"
)

func PrintEven() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func PrintOdd() {
	for i := 0; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func Numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("%d ", i)
	}
}

func Alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c ", i)
	}
	fmt.Println()
}
