package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	nums = append(nums, 6, 7, 8)

	fmt.Println("Slice:", nums)
	fmt.Println("Tamanho:", len(nums))
	fmt.Println("Capacidade:", cap(nums))
}
