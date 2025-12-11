package main

import (
	"fmt"
)

// cubeToFloat64 преобразует uint8 в float64 и возводит в куб.
func cubeToFloat64(v uint8) float64 {
	x := float64(v)
	return x * x * x
}

// CubePipeline читает числа из in, преобразует их и пишет в out.
// Когда входной канал закрывается, функция закрывает выходной канал и завершается.
func CubePipeline(in <-chan uint8, out chan<- float64) {
	for v := range in {
		out <- cubeToFloat64(v)
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	// Запускаем конвейер
	go CubePipeline(in, out)

	// Отправляем несколько чисел в первый канал
	go func() {
		nums := []uint8{1, 2, 3, 5, 10}
		for _, n := range nums {
			in <- n
		}
		close(in) // важно закрыть, чтобы конвейер понял, что данные закончились
	}()

	// Читаем результаты из второго канала и выводим на экран
	for res := range out {
		fmt.Println(res)
	}
}
