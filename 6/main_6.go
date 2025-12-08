package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGeneratorWithRand(count, min, max int, r *rand.Rand) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			n := r.Intn(max-min+1) + min
			ch <- n
		}
	}()
	return ch
}

func RandomGenerator(count, min, max int) <-chan int {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randomGeneratorWithRand(count, min, max, src)
}

func main() {
	ch := RandomGenerator(5, 1, 10)
	for v := range ch {
		fmt.Println(v)
	}
}
