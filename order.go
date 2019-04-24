package main

import "fmt"

// Order (reverse) an Array (Slice in Go)
// Ex:
//   []{1,2,3,4,5,6,7,8,9,10}
// Result:
//   []{10,9,8,7,6,5,4,3,2,1}

// Track the first and last index
// In a loop swap both of them
// then: start = start + 1 and end = end - 1

func (a *Array) reverseSwappingStartEnd(start, end int) error {
	for start < end {
		a.dst[start], a.dst[end] = a.dst[end], a.dst[start]

		start++
		end--
	}

	return nil
}

func (a *Array) reverseSwappingStartEndRecursive(start, end int) {
	if start >= end {
		return
	}

	a.dstGeneric[start], a.dstGeneric[end] = a.dstGeneric[end], a.dstGeneric[start]

	a.reverseSwappingStartEndRecursive(start+1, end-1)
}

func (a *Array) runAndPrintreverseSwappingStartEnd() error {
	copy(a.dst, input)
	copy(a.dstGeneric, input)

	fmt.Println(printWithColors("Reverse Array Swapping Start and End", 4))
	fmt.Printf("Original array: %s Length: %s\n", printWithColors(a.src, 3), printWithColors(len(a.src), 2))

	if err := a.reverseSwappingStartEnd(0, len(a.dst)-1); err != nil {
		return err
	}

	fmt.Printf("Iterative way: %s Length: %s\n", printWithColors(a.dst, 3), printWithColors(len(a.dst), 2))

	a.reverseSwappingStartEndRecursive(0, len(a.dstGeneric)-1)

	fmt.Printf("Recursive way: %s Length: %s\n", printWithColors(a.dstGeneric, 3), printWithColors(len(a.dstGeneric), 2))

	return nil
}

// Order (sorted) an Array (Slice in Go)
// Ex:
//   []{10,8,1,3,7,6,9,5,4,2}
// Result:
//   []{1,2,3,4,5,6,7,8,9,10}

// Iterate from
// In a loop swap both of them
// then: start = start + 1 and end = end - 1

func (a *Array) bubbleSort() error {
	// Iterate the array from end to start
	for i := range a.dst {
		i = len(a.dst) - 1 - i

		// Iterate from the second element to the element with the index i
		for j := 1; j <= i; j++ {
			if a.dst[j-1] > a.dst[j] {
				// Change the order of the elements if the left with minor index is
				// greater than the other
				a.dst[j-1], a.dst[j] = a.dst[j], a.dst[j-1]
			}
		}
	}

	return nil
}

func (a *Array) runAndPrintbubbleSort() error {
	copy(a.src, inputUnordered)
	copy(a.dst, inputUnordered)
	copy(a.dstGeneric, inputUnordered)

	fmt.Println(printWithColors("Bubble Sort", 4))
	fmt.Printf("Original array: %s Length: %s\n", printWithColors(a.src, 2), printWithColors(len(a.src), 3))

	if err := a.bubbleSort(); err != nil {
		return err
	}

	fmt.Printf("Sorted: %s Length: %s\n", printWithColors(a.dst, 2), printWithColors(len(a.dst), 3))

	return nil
}
