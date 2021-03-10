/*
Протестируйте производительность операций чтения и записи
на множестве действительных чисел,
безопасность которого обеспечивается sync.Mutex
для разных вариантов использования:
50% запись, 50% чтение;
 */
package main

import (
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
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
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

func BenchmarkSetAdd(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(500)
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
		b.SetParallelism(500)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}
