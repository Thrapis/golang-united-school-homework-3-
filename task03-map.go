package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, 0, len(input))
	var indexes = make([]int, 0, len(input))
	for i, _ := range input {
		indexes = append(indexes, i)
	}
	sort.Ints(indexes)
	for _, vi := range indexes {
		result = append(result, input[vi])
	}
	return
}
