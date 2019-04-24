package main

var input = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
var inputUnordered = []int{10, 3, 9, 2, 6, 4, 0, 1, 7, 5, 8}

type Array struct {
	dst        []int
	dstGeneric []int
	elements   int
	src        []int
}

func main() {
	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	a := Array{
		src:        input,
		elements:   5,
		dst:        inputCopy,
		dstGeneric: inputCopy,
	}

	if err := a.runAndPrintLeftRotate(); err != nil {
		panic(err)
	}

	if err := a.runAndPrintreverseSwappingStartEnd(); err != nil {
		panic(err)
	}

	if err := a.runAndPrintbubbleSort(); err != nil {
		panic(err)
	}
}
