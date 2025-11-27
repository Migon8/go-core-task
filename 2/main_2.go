package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sliceExample возвращает новый слайс, содержащий только чётные числа
func sliceExample(nums []int) []int {
	evens := make([]int, 0, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}

// addElements добавляет число value в конец слайса и возвращает новый слайс
func addElements(nums []int, value int) []int {
	// можно сразу append к nums, это всё равно создаст новый/расширенный слайс
	return append(nums, value)
}

// copySlice возвращает полную копию слайса
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

// helper для генерации случайного слайса из 10 элементов
func generateRandomSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		// допустим, случайные числа от 0 до 99
		nums[i] = rand.Intn(100)
	}
	return nums
}

func main() {
	// 1. Создаём originalSlice из случайных значений
	originalSlice := generateRandomSlice(10)
	fmt.Println("Original slice:", originalSlice)

	// 2. sliceExample — только чётные
	evens := sliceExample(originalSlice)
	fmt.Println("Even numbers:", evens)

	// 3. addElements — добавляем число
	withAdded := addElements(originalSlice, 42)
	fmt.Println("After addElements(42):", withAdded)

	// 4. copySlice — проверяем независимость копии
	copied := copySlice(originalSlice)
	fmt.Println("Copied slice:", copied)

	// изменяем оригинал, чтобы проверить, что копия не меняется
	if len(originalSlice) > 0 {
		originalSlice[0] = -1
	}
	fmt.Println("Original after modification:", originalSlice)
	fmt.Println("Copied after original modification:", copied)

	// 5. removeElement — удаляем элемент по индексу
	if len(originalSlice) > 3 {
		removed := removeElement(originalSlice, 3)
		fmt.Println("After removeElement index 3:", removed)
	} else {
		fmt.Println("Slice too short to remove index 3")
	}
}
