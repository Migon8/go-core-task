package main

import (
	"reflect"
	"testing"
)

func TestCubeToFloat64(t *testing.T) {
	tests := []struct {
		name string
		in   uint8
		want float64
	}{
		{"zero", 0, 0},
		{"one", 1, 1},
		{"two", 2, 8},
		{"five", 5, 125},
		{"ten", 10, 1000},
		{"maxUint8", 255, float64(255 * 255 * 255)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cubeToFloat64(tt.in)
			if got != tt.want {
				t.Fatalf("cubeToFloat64(%d) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	// Запускаем конвейер
	go CubePipeline(in, out)

	// Входные данные
	input := []uint8{0, 1, 2, 5, 10, 255}
	expected := []float64{
		0,
		1,
		8,
		125,
		1000,
		float64(255 * 255 * 255),
	}

	// Пишем в входной канал в отдельной горутине
	go func() {
		for _, v := range input {
			in <- v
		}
		close(in)
	}()

	// Читаем из выхода
	var result []float64
	for v := range out {
		result = append(result, v)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("unexpected pipeline result: got %v, want %v", result, expected)
	}
}
