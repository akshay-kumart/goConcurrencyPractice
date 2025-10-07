package concurrency

func digits(number int, dchn1 chan int) {
	for number != 0 {
		digit := number % 10
		dchn1 <- digit
		number /= 10
	}
	close(dchn1)
}

func CalcSquares(number int, sqaureCh chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	sqaureCh <- sum
}
func CalcCubes(number int, cubeCh chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeCh <- sum
}
