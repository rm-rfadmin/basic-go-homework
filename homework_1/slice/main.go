package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := DeleteSliceIndex(s1, 2)
	fmt.Println(s2)

}