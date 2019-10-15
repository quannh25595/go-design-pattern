package structural

import "log"

type Function func(int) int

func LogDecorate(fn Function) Function {
	return func(n int) int {
		log.Println("Starting execution with the integer ", n)
		result := fn(n)
		log.Println("Result ", result)
		return result
	}
}

//Usage

func Double(n int) int {
	return n * 2
}

func decorator() {
	fn := LogDecorate(Double)
	fn(5)
}
