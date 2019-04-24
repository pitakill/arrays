package main

import "fmt"

// Rotate an Array (Slice in Go) by n elements
// Ex:
//   []{1,2,3,4,5,6,7,8,9,10}
//   by 3 elements
// Result:
//   []{4,5,6,7,8,9,10,1,2,3}

// Store n first elements to temp array
// Shift elements of array
// Store back the n elements
func (a *Array) leftRotate(n int) error {
	a.dst = append(a.dst[n:], append(a.dst[0:n])...)

	return nil
}
func (a *Array) leftRotate_Generic(n int) error {
	temp := make([]int, n)

	for i := 0; i < n; i++ {
		temp[i] = a.dstGeneric[i]
	}
	for i := n; i < len(a.dstGeneric); i++ {
		a.dstGeneric[i-n] = a.dstGeneric[i]
	}
	for i := 0; i < len(temp); i++ {
		a.dstGeneric[len(a.dstGeneric)-n+i] = temp[i]
	}

	return nil
}

func (a *Array) runAndPrintLeftRotate() error {
	fmt.Println(printWithColors("Rotate Left", 4))
	fmt.Printf("Orginal array: %s rotate %s elements. Length: %s\n", printWithColors(a.src, 2), printWithColors(a.elements, 4), printWithColors(len(a.src), 3))

	if err := a.leftRotate(a.elements); err != nil {
		return err
	}
	fmt.Printf("Result with built-in functions: %s Length: %s\n", printWithColors(a.dst, 2), printWithColors(len(a.dst), 3))

	if err := a.leftRotate_Generic(a.elements); err != nil {
		return err
	}
	fmt.Printf("Result with generic for: %s Length: %s\n", printWithColors(a.dstGeneric, 2), printWithColors(len(a.dstGeneric), 3))

	return nil
}
