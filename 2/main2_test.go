package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"only odd", []int{1, 3, 5}, []int{}},
		{"only even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceExample(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	nums := []int{1, 2, 3}
	got := addElements(nums, 4)
	want := []int{1, 2, 3, 4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements(%v, 4) = %v, want %v", nums, got, want)
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3}
	copied := copySlice(original)

	if !reflect.DeepEqual(copied, original) {
		t.Fatalf("copySlice(%v) = %v, want same values", original, copied)
	}

	// проверяем, что это другая память (изменения не протекают)
	original[0] = 99
	if copied[0] == original[0] {
		t.Errorf("copied slice is affected by changes in original: copied=%v, original=%v", copied, original)
	}
}

func TestRemoveElement(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}

	t.Run("remove middle", func(t *testing.T) {
		got := removeElement(nums, 2) // удаляем 30
		want := []int{10, 20, 40, 50}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("removeElement(%v, 2) = %v, want %v", nums, got, want)
		}
	})

	t.Run("remove first", func(t *testing.T) {
		got := removeElement(nums, 0)
		want := []int{20, 30, 40, 50}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("removeElement(%v, 0) = %v, want %v", nums, got, want)
		}
	})

	t.Run("remove last", func(t *testing.T) {
		got := removeElement(nums, len(nums)-1)
		want := []int{10, 20, 30, 40}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("removeElement(%v, last) = %v, want %v", nums, got, want)
		}
	})

	t.Run("invalid index", func(t *testing.T) {
		got := removeElement(nums, 100)
		want := nums // при некорректном индексе возвращаем копию, но по значениям совпадает
		if !reflect.DeepEqual(got, want) {
			t.Errorf("removeElement(%v, 100) = %v, want %v", nums, got, want)
		}
		if &got[0] == &nums[0] {
			t.Errorf("expected new slice on invalid index, but got same underlying array")
		}
	})
}
