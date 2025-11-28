package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(nums []int) []int {
	evens := make([]int, 0, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}

func addElements(nums []int, value int) []int {

	return append(nums, value)
}

func copySlice(nums []int) []int {
	if nums == nil {
		return nil
	}
	cp := make([]int, len(nums))
	copy(cp, nums)
	return cp
}

// removeElement удаляет элемент по индексу и возвращает новый слайс без него
// Если индекс некорректный — возвращаем копию исходного слайса без изменений.

func removeElement(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return copySlice(nums)
	}

	result := make([]int, 0, len(nums)-1)
	result = append(result, nums[:index]...)
	result = append(result, nums[index+1:]...)
	return result
}

func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		// случайные числа от 0 до 99
		nums[i] = rand.Intn(100)
	}
	return nums
}

func main() {

	originalSlice := generateRandomSlice(10)
	fmt.Println("Original slice:", originalSlice)

	evens := sliceExample(originalSlice)
	fmt.Println("Even numbers:", evens)

	withAdded := addElements(originalSlice, 42)
	fmt.Println("After addElements(42):", withAdded)

	copied := copySlice(originalSlice)
	fmt.Println("Copied slice:", copied)

	// изменяем оригинал, чтобы проверить, что копия не меняется
	if len(originalSlice) > 0 {
		originalSlice[0] = -1
	}
	fmt.Println("Original after modification:", originalSlice)
	fmt.Println("Copied after original modification:", copied)

	//  removeElement — удаляем элемент по индексу
	if len(originalSlice) > 3 {
		removed := removeElement(originalSlice, 3)
		fmt.Println("After removeElement index 3:", removed)
	} else {
		fmt.Println("Slice too short to remove index 3")
	}
}
