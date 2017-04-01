package mutexmultimap

import (
	"sync"
)

// MutexMultiMap that maps keys to values, similar to Map
// but in which each key may be associated with multiple values.
// a → 1, 2
// b → 3
type MutexMultiMap struct {
	sync.RWMutex
	hint  int
	items map[string][]interface{}
}

// New is sole constructor.
func New(hint int) *MutexMultiMap {
	return &MutexMultiMap{
		hint:  hint,
		items: make(map[string][]interface{}, hint),
	}
}

// Clear is removes all key-value pairs from the multimap, leaving it empty.
func (m *MutexMultiMap) Clear() {
	m.Lock()
	m.items = make(map[string][]interface{}, m.hint)
	m.Unlock()
}

// ContainsEntry is returns true if this multimap contains
// at least one key-value pair with the key key and the value value.
func (m *MutexMultiMap) ContainsEntry(key string, value interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	items, ok := m.items[key]
	if !ok {
		return false
	}
	for i := range items {
		if items[i] == value {
			return true
		}
	}
	return false
}

// ContainsKey is returns true if this multimap contains at least one key-value
// pair with the key key.
func (m *MutexMultiMap) ContainsKey(key string) bool {
	m.RLock()
	_, isExist := m.items[key]
	m.RUnlock()
	return isExist
}

// ContainsValue is returns true if this multimap contains at least one
// key-value pair with the value value.
func (m *MutexMultiMap) ContainsValue(value interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	for _, items := range m.items {
		for i := range items {
			if items[i] == value {
				return true
			}
		}
	}
	return false
}

// Entries is returns a view collection of all key-value pairs contained in
// this multimap, as Map.Entry instances.
func (m *MutexMultiMap) Entries() map[string][]interface{} {
	m.RLock()
	src := m.items
	dst := make(map[string][]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	m.RUnlock()
	return dst
}

// Equals is compares the specified object with this multimap for equality.
func (m *MutexMultiMap) Equals(src map[string][]interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	if len(m.items) != len(src) {
		return false
	}
	for k, items := range m.items {
		for i := range items {
			if m.items[k][i] != src[k][i] {
				return false
			}
		}
	}
	return true
}

// Get is returns a view collection of the values associated with key in
// this multimap, if any.
func (m *MutexMultiMap) Get(key string) []interface{} {
	m.RLock()
	src := m.items[key]
	dst := make([]interface{}, len(src))
	for i := range src {
		dst[i] = src[i]
	}
	m.RUnlock()
	return dst
}

// IsEmpty is returns true if this multimap contains no key-value pairs.
func (m *MutexMultiMap) IsEmpty() bool {
	return m.Len() == 0
}

// Put is stores a key-value pair in this multimap.
func (m *MutexMultiMap) Put(key string, value interface{}) {
	m.Lock()
	items := m.items[key]
	items = append(items, value)
	m.items[key] = items
	m.Unlock()
}

// PutAll is stores a key-value pair in this multimap for each of values
// all using the same key, key.
func (m *MutexMultiMap) PutAll(key string, values []interface{}) {
	m.Lock()
	m.items[key] = values
	m.Unlock()
}

// Delete is removes a single key-value pair with the key key and
// the value value from this multimap, if such exists.
func (m *MutexMultiMap) Delete(key string, value interface{}) {
	m.Lock()
	defer m.Unlock()
	items, ok := m.items[key]
	if !ok {
		return
	}
	cutSlice(&items, value)
	if len(items) == 0 {
		delete(m.items, key)
		return
	}
	m.items[key] = items
}

// DeleteAll is removes all values associated with the key key.
func (m *MutexMultiMap) DeleteAll(key string) []interface{} {
	m.Lock()
	defer m.Unlock()
	items, ok := m.items[key]
	if !ok {
		return nil
	}
	m.items[key] = make([]interface{}, 0, len(items))
	return items
}

// ReplaceValues is stores a collection of values with the same key,
// replacing any existing values for that key.
func (m *MutexMultiMap) ReplaceValues(key string, srcs ...interface{}) []interface{} {
	m.Lock()
	defer m.Unlock()
	_, ok := m.items[key]
	if !ok {
		return nil
	}
	m.items[key] = srcs
	return m.items[key]
}

// Len is returns the number of key-value pairs in this multimap.
func (m *MutexMultiMap) Len() int {
	m.RLock()
	l := len(m.items)
	m.RUnlock()
	return l
}

func cutSlice(ps *[]interface{}, value interface{}) {
	items := *ps
	for i := 0; i < len(items); {
		if items[i] == value {
			copy(items[i:], items[i+1:])
			items = items[:len(items)-1]
			continue
		}
		i++
	}
	*ps = items
}
