package main

import "testing"

func TestAddAndGet(t *testing.T) {
	m := &StringIntMap{}

	m.Add("foo", 10)

	value, ok := m.Get("foo")
	if !ok {
		t.Fatalf("ожидали, что ключ 'foo' существует")
	}

	if value != 10 {
		t.Fatalf("ожидали значение 10, получили %d", value)
	}
}

func TestExists(t *testing.T) {
	m := &StringIntMap{}

	m.Add("bar", 20)

	if !m.Exists("bar") {
		t.Fatalf("ожидали, что ключ 'bar' существует")
	}

	if m.Exists("baz") {
		t.Fatalf("ожидали, что ключ 'baz' не существует")
	}
}

func TestRemove(t *testing.T) {
	m := &StringIntMap{}

	m.Add("key", 100)
	m.Remove("key")

	if m.Exists("key") {
		t.Fatalf("после Remove ключ 'key' не должен существовать")
	}

	_, ok := m.Get("key")
	if ok {
		t.Fatalf("после Remove Get должен возвращать ok = false")
	}
}

func TestCopy(t *testing.T) {
	m := &StringIntMap{}
	m.Add("a", 1)
	m.Add("b", 2)

	copyMap := m.Copy()

	// Проверяем, что в копии есть те же значения
	if len(copyMap) != 2 {
		t.Fatalf("в копии ожидали 2 элемента, получили %d", len(copyMap))
	}

	if copyMap["a"] != 1 || copyMap["b"] != 2 {
		t.Fatalf("ожидали значения a=1, b=2 в копии, получили: %#v", copyMap)
	}

	// Меняем оригинал и смотрим, что копия не изменилась
	m.Add("c", 3)
	if len(copyMap) != 2 {
		t.Fatalf("копия не должна меняться при изменении оригинала, а длина стала %d", len(copyMap))
	}
}
