package main

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name          string
		a             []int
		b             []int
		wantHas       bool
		wantIntersect []int
	}{
		{
			name:          "example from task",
			a:             []int{65, 3, 58, 678, 64},
			b:             []int{64, 2, 3, 43},
			wantHas:       true,
			wantIntersect: []int{64, 3},
		},
		{
			name:          "no intersection",
			a:             []int{1, 2, 3},
			b:             []int{4, 5, 6},
			wantHas:       false,
			wantIntersect: []int{},
		},
		{
			name:          "one slice is empty",
			a:             []int{},
			b:             []int{1, 2, 3},
			wantHas:       false,
			wantIntersect: []int{},
		},
		{
			name:          "both slices empty",
			a:             []int{},
			b:             []int{},
			wantHas:       false,
			wantIntersect: []int{},
		},
		{
			name:    "duplicates in both slices",
			a:       []int{1, 1, 2, 2, 3},
			b:       []int{3, 3, 2, 1, 1},
			wantHas: true,

			wantIntersect: []int{3, 2, 1},
		},
		{
			name:          "intersection only one value",
			a:             []int{10, 20, 30},
			b:             []int{5, 20, 40},
			wantHas:       true,
			wantIntersect: []int{20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHas, gotIntersect := Intersect(tt.a, tt.b)

			if gotHas != tt.wantHas {
				t.Errorf("Intersect() has = %v, want %v", gotHas, tt.wantHas)
			}

			if !reflect.DeepEqual(gotIntersect, tt.wantIntersect) {
				t.Errorf("Intersect() intersect = %v, want %v", gotIntersect, tt.wantIntersect)
			}
		})
	}
}
