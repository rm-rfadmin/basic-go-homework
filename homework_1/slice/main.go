package main

import "fmt"

func DeleteSliceIndex(slice []int, index int) []int {
	s1 := slice[:index]
	s2 := slice[index+1:]
	return append(s1, s2...)
}

func DeleteSliceIndex2(slice []int, index int) []int {

	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]

}

func DeleteSliceIndex3(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}

func CheckSlice[T any](slice []T, index int) bool {
	length := len(slice)
	if index < 0 || index >= length {
		return false
	}
	return true
}

// func DeleteSliceIndex4[T []int | []uint](slice T, index int) T {  ????? []int | []uint 为什么不行？？？
func DeleteSliceIndex4[T any](slice []T, index int) ([]T, error) {
	if !CheckSlice(slice, index) {
		return nil, fmt.Errorf("下标超出范围，长度 %d, 下标 %d", len(slice), index)
	} else {
		s1 := slice[:index]
		s2 := slice[index+1:]
		s3 := append(s1, s2...)
		return s3, nil
	}
}

func DeleteSliceIndex5[T any](slice []T, index int) []T {
	if !CheckSlice(slice, index) {
		return nil
	} else {
		copy(slice[index:], slice[index+1:])

		newLen := len(slice) - 1
		slice = slice[:len(slice)-1]

		if newLen < cap(slice)/2 {
			newSlice := make([]T, newLen) // TODO：缩容后的容量应该比长度大，否则会重新分配内存
			copy(newSlice, slice)
			slice = newSlice
		}
		return slice
	}
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := DeleteSliceIndex(s1, 2)
	fmt.Println(s2)

	s3 := []int{1, 2, 3, 4, 5}

	s4 := DeleteSliceIndex2(s3, 2)
	fmt.Println(s4)

	s5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p", &s5)
	s6 := DeleteSliceIndex3(s5, 2)
	fmt.Printf("%p", &s6)
	fmt.Println(s6)

	// 泛型版本
	s7 := []int{1, 2, 3, 4, 5}
	s8, err := DeleteSliceIndex4(s7, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s8)
	}

	s9 := []int{1, 2, 3, 4, 5}
	s10 := DeleteSliceIndex5(s9, 2)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = DeleteSliceIndex5(s10, 2)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = DeleteSliceIndex5(s10, 2)
	fmt.Println(s10, len(s10), cap(s10))
	s10 = DeleteSliceIndex5(s10, 0)
	fmt.Println(s10, len(s10), cap(s10))

}
