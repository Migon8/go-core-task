package main

func Intersect(a, b []int) (bool, []int) {
	// Кладём все элементы из a в map для быстрых проверок
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}

	result := make([]int, 0)
	seen := make(map[int]struct{}) // чтобы не добавлять одно и то же число несколько раз

	// Идём по b и смотрим, какие элементы есть в m
	for _, v := range b {
		if _, ok := m[v]; ok {
			// элемент есть в a, проверяем, не добавляли ли мы его уже
			if _, alreadyAdded := seen[v]; !alreadyAdded {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}

	return len(result) > 0, result
}
