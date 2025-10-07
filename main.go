package main

import (
	"fmt"
	"goPractice/concurrency"
	"sync"
)

func main() {
	//fmt.Println("Heeloo")
	//go concurrency.PrintEven()
	//go concurrency.PrintOdd()
	//time.Sleep(time.Second * 3)

	go concurrency.Numbers()
	go concurrency.Alphabets()
	//time.Sleep(time.Millisecond * 4000)
	//fmt.Println("terminatedd")
	//------------------------------------
	//concurrency.NilChannel()
	done := make(chan bool)
	go concurrency.HelloChannel(done) //other way to run go func without time.Sleep
	<-done
	//fmt.Println("main func")
	//----------------------------------------------
	// when channels are used then no need to use time.Sleep or waitGroups
	resultNumSq := make(chan int)
	resultNumCu := make(chan int)
	go concurrency.CaluclateSquares(2, resultNumSq)
	go concurrency.CaluclateCubes(2, resultNumCu)
	//fmt.Println(<-resultNumSq)
	//fmt.Println(<-resultNumCu)

	//----------------------------------- Int Ques

	sumSqChan := make(chan int)
	go concurrency.ChannelSumOfNSquare(10, sumSqChan)
	//fmt.Println(<-sumSqChan)

	//---------------------------------------------------- Int ques
	for i := 0; i < 5; i++ {
		//go concurrency.ChannelSumOfNSquare(i, sumSqChan)
		//fmt.Println(<-sumSqChan)
	}
	//----------------------------------------------------
	// sendCh chan<- int - send only
	// receiveCh <-chan int - receive only

	sendCh := make(chan int)
	receiveCh := make(chan int)

	//go concurrency.SendChan(sendCh)
	//fmt.Println(<-sendCh)
	//receiveCh <- 50
	//go concurrency.ReceiveChan(receiveCh)
	//----------------------------------- above cause deadlock

	// Start receiver BEFORE sending to receiveCh
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		concurrency.ReceiveChan(receiveCh)
	}()

	// Separate: send/receive via sendCh
	go concurrency.SendChan(sendCh)
	fmt.Println(<-sendCh) // prints 1000

	receiveCh <- 50 // safe: receiver is already waiting
	wg.Wait()       // ensure receiver prints before main exits
	fmt.Println("-----------------")

	//---------------------------------------------------loops on channel

	chn1 := make(chan int)
	go concurrency.Producer(chn1)
	for {
		v, ok := <-chn1 //if ok is false then channel is closed
		if !ok {
			break
		}
		fmt.Print("Received ", v, ok)
	}
	fmt.Println()
	fmt.Println("-----------------------------")
	//--------------------------------------new way of finding sum of sq using ch

	sqrch := make(chan int)
	cubech := make(chan int)

	go concurrency.CalcSquares(4, sqrch)
	go concurrency.CalcCubes(4, cubech)

	fmt.Println(<-sqrch)
	fmt.Println(<-cubech)
	// fmt.Println("Sum of square and cube ", <-sqrch+<-cubech) we should not do like this because of deadlock

	//a, b := <-sqrch, <-cubech 			// need to fix
	//fmt.Println("Sum of square and cube is = ", a+b)
}
