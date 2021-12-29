package cmap

import (
	"constraints"
	"sync"
)

type Map[K constraints.Ordered, V any] interface {
	Clear() Map[K, V]
	Add(k K, v V) Map[K, V]
	AddAll(m map[K]V) Map[K, V]
	Delete(k K) Map[K, V]
	Immutable() ImmutableMap[K, V]

	Len() int
	Get(k K) (V, bool)
	Contains(k K) bool
	Keys() []K
	Values() []V
	Map() map[K]V
	Range(apply func(a K, v V))
}

type ImmutableMap[K constraints.Ordered, V any] interface {
	Get(k K) (V, bool)
	Contains(k K) bool
	Keys() []K
	Values() []V
	Map() map[K]V
	Range(apply func(a K, v V))
}

func New[K constraints.Ordered, V any]() Map[K, V] {
	return &cmap[K, V]{
		store: map[K]V{},
	}
}

func From[K constraints.Ordered, V any](m map[K]V) Map[K, V] {
	return New[K, V]().AddAll(m)
}

type cmap[K constraints.Ordered, V any] struct {
	store     map[K]V
	mu        sync.RWMutex
	immutable bool
}

func (m *cmap[K, V]) Clear() Map[K, V] {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.immutable {
		panic("can no modify a immutable map")
	}

	m.store = map[K]V{}
	return m
}
func (m *cmap[K, V]) Add(k K, v V) Map[K, V] {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.immutable {
		panic("can no modify a immutable map")
	}

	m.store[k] = v
	return m
}
func (m *cmap[K, V]) AddAll(all map[K]V) Map[K, V] {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.immutable {
		panic("can no modify a immutable map")
	}

	for k, v := range all {
		m.store[k] = v
	}
	return m
}

func (m *cmap[K, V]) Delete(k K) Map[K, V] {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.immutable {
		panic("can no modify a immutable map")
	}

	delete(m.store, k)
	return m
}

func (m *cmap[K, V]) Get(k K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	v, ok := m.store[k]
	return v, ok
}
func (m *cmap[K, V]) Contains(k K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, exist := m.store[k]
	return exist
}

func (m *cmap[K, V]) Keys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var keys []K
	for k := range m.store {
		keys = append(keys, k)
	}
	return keys
}
func (m *cmap[K, V]) Values() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var values []V
	for _, v := range m.store {
		values = append(values, v)
	}
	return values
}
func (m *cmap[K, V]) Map() map[K]V {
	m.mu.RLock()
	defer m.mu.RUnlock()

	ma := map[K]V{}
	for k, v := range m.store {
		ma[k] = v
	}
	return ma
}
func (m *cmap[K, V]) Range(apply func(a K, v V)) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for k, v := range m.store {
		apply(k, v)
	}
}

func (m *cmap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.store)
}

func (m *cmap[K, V]) Immutable() ImmutableMap[K, V] {
	m.mu.RLock()
	defer m.mu.RUnlock()

	m.immutable = true
	return m
}
