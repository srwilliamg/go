package slice

import (
	"fmt"
	"sync"
)

func RunSliceTest(wg *sync.WaitGroup) {
	defer wg.Done()
	// Create a slice of strings.
	// Contains a length and capacity of 5 elements.
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// Slice the third element and restrict the capacity.
	// Contains a length and capacity of 1 element.
	slice := source[2:3:3]
	// Append a new string to the slice.
	slice = append(slice, "Kiwi")

	fmt.Println("source: ", source)
	fmt.Println("slice: ", slice)

	source = append(source, "Green_Apple")
	slice = append(slice, "Red_Apple")
	fmt.Println("source: ", source)
	fmt.Println("slice: ", slice)

	source2 := make([]string, 3, 6)
	bigSlice := append(source2, source...)
	fmt.Println("bigSlice: ", bigSlice)
	source2 = append(source2, "BIGFRUIT")
	fmt.Println("source: ", source)
	fmt.Println("source2: ", source2)
}
