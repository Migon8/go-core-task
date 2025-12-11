package merge

import (
	"context"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMergeBasic(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 10
	}()

	go func() {
		defer close(ch3)
		ch3 <- 100
		ch3 <- 200
	}()

	out := Merge[int](ctx, ch1, ch2, ch3)

	var got []int
	for v := range out {
		got = append(got, v)
	}

	sort.Ints(got)
	want := []int{1, 2, 10, 100, 200}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected result: got %v, want %v", got, want)
	}
}

func TestMergeNoChannels(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := Merge[int](ctx)

	select {
	case _, ok := <-out:
		if ok {
			t.Fatalf("expected closed channel, but got value")
		}
	case <-time.After(time.Second):
		t.Fatalf("timeout: channel not closed")
	}
}

func TestMergeContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)
	out := Merge[int](ctx, ch)

	time.AfterFunc(100*time.Millisecond, func() {
		cancel()
		close(ch)
	})

	select {
	case <-out:
	case <-time.After(2 * time.Second):
		t.Fatalf("timeout: merge did not stop after context cancel")
	}
}
