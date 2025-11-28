package main

// StringIntMap — обёртка над map[string]int.
type StringIntMap struct {
	data map[string]int
}

// Add добавляет или обновляет значение по ключу.
func (m *StringIntMap) Add(key string, value int) {
	// Ленивое создание мапы, если она ещё nil.
	if m.data == nil {
		m.data = make(map[string]int)
	}
	m.data[key] = value
}

// Remove удаляет элемент по ключу.
func (m *StringIntMap) Remove(key string) {
	if m.data == nil {
		return
	}
	delete(m.data, key)
}

// Copy возвращает новую map[string]int с копией всех элементов.
func (m *StringIntMap) Copy() map[string]int {
	result := make(map[string]int)

	if m.data == nil {
		return result
	}

	for k, v := range m.data {
		result[k] = v
	}

	return result
}

// Exists проверяет наличие ключа в карте.
func (m *StringIntMap) Exists(key string) bool {
	if m.data == nil {
		return false
	}

	_, ok := m.data[key]
	return ok
}

// Get возвращает значение и флаг, найден ли ключ.
func (m *StringIntMap) Get(key string) (int, bool) {
	if m.data == nil {
		return 0, false
	}

	v, ok := m.data[key]
	return v, ok
}
