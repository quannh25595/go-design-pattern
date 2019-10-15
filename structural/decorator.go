package structural

import "log"

type Object func(int) int

func LogDecorate(fn Object) Object {
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
