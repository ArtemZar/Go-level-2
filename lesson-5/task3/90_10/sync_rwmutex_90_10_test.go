/*
Протестируйте производительность операций чтения и записи
на множестве действительных чисел,
безопасность которого обеспечивается sync.RWMutex
для разных вариантов использования:
90% запись, 10% чтение
*/
package main


import (
	"sync"
	"testing"
)

type Set struct {
	sync.RWMutex
	mm map[int]struct{}
}


func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func(s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func(s *Set) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}


func BenchmarkSetAdd(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(900)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}


func BenchmarkSetHas(b *testing.B) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(100)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}
