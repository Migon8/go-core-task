package main

import (
	"sync/atomic"
	"testing"
	"time"
)

// Проверяем, что Wait действительно ждёт всех горутин.
func TestSemaphoreWaitGroup_WaitAll(t *testing.T) {
	const workers = 10

	wg := NewSemaphoreWaitGroup(workers)

	var counter int64

	for i := 0; i < workers; i++ {
		go func() {
			// имитируем какую-то работу
			time.Sleep(10 * time.Millisecond)

			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	if counter != workers {
		t.Fatalf("ожидали counter = %d, получили %d", workers, counter)
	}
}

// Проверяем, что для 0 горутин Wait не падает в дедлок и
// отрабатывает мгновенно.
func TestSemaphoreWaitGroup_ZeroWorkers(t *testing.T) {
	wg := NewSemaphoreWaitGroup(0)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// всё ок, Wait вернулся
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait заблокировался для 0 горутин")
	}
}
