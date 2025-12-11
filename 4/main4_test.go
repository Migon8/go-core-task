package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestDifference_EmptySecond(t *testing.T) {
	slice1 := []string{"a", "b", "c"}
	slice2 := []string{}
	expected := []string{"a", "b", "c"}

	result := Difference(slice1, slice2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestDifference_EmptyFirst(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"a", "b"}

	result := Difference(slice1, slice2)

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func TestDifference_NoDifference(t *testing.T) {
	slice1 := []string{"a", "b"}
	slice2 := []string{"a", "b"}

	result := Difference(slice1, slice2)

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}
