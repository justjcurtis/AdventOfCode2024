package utils

import "sync"

func SynMapLength(m *sync.Map) int {
	length := 0
	m.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}
