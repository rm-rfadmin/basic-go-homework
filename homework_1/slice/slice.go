package main

func DeleteSliceIndex(slice []int, index int) []int {
	s1 := slice[:index]
	s2 := slice[index+1:]
	return append(s1, s2...)
}
