package encoder

import (
	"os"
	"sort"
)

func FrequencySlice(file *os.File) [][]interface{} {
	m := runeCounter(file)

	kvPairs := [][]interface{}{}
	for key, value := range m.Data {
		kvPairs = append(kvPairs, []interface{}{key, value})
	}

	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i][1].(int) > kvPairs[j][1].(int)
	})

	return kvPairs
}
