package main

import "fmt"

type FibNumberHistory struct {
	SecondLastNum int
	LastNum       int
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibHist := FibNumberHistory{0, 0}
	fibCallCounter := 0
	return func() int {
		if fibCallCounter == 0 {
			fibCallCounter++
			return fibHist.LastNum
		} else if fibCallCounter == 1 {
			fibCallCounter++
			fibHist.LastNum = 1
			return fibHist.LastNum
		} else {
			currentFib := fibHist.SecondLastNum + fibHist.LastNum
			fibHist.SecondLastNum = fibHist.LastNum
			fibHist.LastNum = currentFib
			return currentFib
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
