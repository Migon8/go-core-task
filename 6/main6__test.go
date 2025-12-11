package main

import (
	"math/rand"
	"testing"
)

func TestRandomGeneratorRange(t *testing.T) {
	count := 100
	min, max := 10, 20
	r := rand.New(rand.NewSource(1))

	ch := randomGeneratorWithRand(count, min, max, r)

	received := 0
	for v := range ch {
		if v < min || v > max {
			t.Fatalf("значение %d вне диапазона [%d, %d]", v, min, max)
		}
		received++
	}

	if received != count {
		t.Fatalf("ожидалось %d значений, получили %d", count, received)
	}
}

func TestRandomGeneratorDeterministic(t *testing.T) {
	count := 5
	min, max := 0, 100
	seed := int64(42)

	r1 := rand.New(rand.NewSource(seed))
	r2 := rand.New(rand.NewSource(seed))

	ch1 := randomGeneratorWithRand(count, min, max, r1)
	ch2 := randomGeneratorWithRand(count, min, max, r2)

	var s1, s2 []int
	for v := range ch1 {
		s1 = append(s1, v)
	}
	for v := range ch2 {
		s2 = append(s2, v)
	}

	if len(s1) != len(s2) {
		t.Fatalf("разная длина: %d и %d", len(s1), len(s2))
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			t.Fatalf("последовательности различаются: %d != %d", s1[i], s2[i])
		}
	}
}
