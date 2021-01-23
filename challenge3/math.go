package math

func Squares(number int) (sum int) {
	sum = 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	return
}

func Cubes(number int) (sum int) {
	sum = 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	return
}

func SquaresPlusCubes(number int) int {
	return Squares(number) + Cubes(number)
}

func SquaresMT(number int, c chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	c <- sum
}

func CubesMT(number int, c chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	c <- sum
}

func SquaresPlusCubesMT(number int, c chan int) {
	sc := make(chan int)
	cc := make(chan int)
	go SquaresMT(number, sc)
	go CubesMT(number, cc)
	squares, cubes := <-sc, <-cc
	c <- squares + cubes
}
