package main

func Difference(slice1, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice2))
	for _, v := range slice2 {
		m[v] = struct{}{}
	}

	result := make([]string, 0)
	for _, v := range slice1 {
		if _, exists := m[v]; !exists {
			result = append(result, v)
		}
	}

	return result
}
