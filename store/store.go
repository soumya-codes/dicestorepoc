package store

import (
	"time"

	"github.com/dolthub/swiss"
	ds "github.com/soumya-codes/AlgoAndDS/generics/datastrcutures"
)

// Store is a generic key-value Store for different DataStructures
type Store[T ds.DSInterface] struct {
	Store   *swiss.Map[string, T]
	Expires *swiss.Map[*T, uint64]
	// TODO: Check if this can better the response latency with this data-structure
	watchedKeys *swiss.Map[string, int]
}

// NewStore returns a new Store
func NewStore[T ds.DSInterface]() *Store[T] {
	return &Store[T]{
		Store:       swiss.NewMap[string, T](64),
		Expires:     swiss.NewMap[*T, uint64](64),
		watchedKeys: swiss.NewMap[string, int](64),
	}
}

func (s *Store[T]) Put(key string, Value T, expDurationMs int64) {
	Value.UpdateLastAccessedAt()
	s.Store.Put(key, Value)
	s.Expires.Put(&Value, uint64(expDurationMs))
}

func (s *Store[T]) setExpiry(Value *T, expDurationMs int64) {
	s.Expires.Put(Value, uint64(time.Now().UnixMilli())+uint64(expDurationMs))
}

func (s *Store[T]) Get(key string) (T, bool) {
	value, exists := s.Store.Get(key)
	if exists {
		value.UpdateLastAccessedAt()
	}

	return value, exists
}

func (s *Store[T]) Delete(key string) bool {
	if _, exists := s.Store.Get(key); exists {
		s.Store.Delete(key)
		return exists
	}

	return false
}

func (s *Store[T]) GetExpiry(key string) (uint64, bool) {
	if val, exists := s.Store.Get(key); exists {
		return s.Expires.Get(&val)
	}

	return 0, false
}

func (s *Store[T]) SetExpiry(key string, expDurationMs int64) {
	obj, exists := s.Store.Get(key)
	if exists {
		s.Expires.Put(&obj, uint64(time.Now().UnixMilli())+uint64(expDurationMs))
	}
}
