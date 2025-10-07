package concurrency

import "fmt"

func NilChannel() {
	var a chan int // will create nil channel
	if a == nil {
		fmt.Println("Channel is nil , now i will initialize it")
		a = make(chan int) //  will initialize unbuffered channel and value will be like '0xc0000200e0'
		fmt.Printf("Type of channel is %T\n", a)
		fmt.Println(a)
	}
}

func HelloChannel(done chan bool) {
	fmt.Println("Hello from Channel")
	done <- true
}

func CaluclateSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func CaluclateCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func ChannelSumOfNSquare(number int, sumSq chan int) {
	result := 0
	for i := 1; i <= number; i++ {
		result += i * i
	}
	sumSq <- result
}

func SendChan(amount chan<- int) {
	amount <- 1000
}

func ReceiveChan(amount <-chan int) {
	fmt.Println(<-amount)
}

func Producer(chn1 chan int) {
	for i := 0; i < 10; i++ {
		chn1 <- i
	}
	close(chn1) //if we dont close channel then it cause deadlock because its keeps on loop in main
}
