package mutexmap

import (
	"sync"
)

// MutexMap is implement mutex map
type MutexMap struct {
	sync.RWMutex
	hint  int
	items map[string]interface{}
}

// New is sole constructor.
func New(hint int) *MutexMap {
	return &MutexMap{
		hint:  hint,
		items: make(map[string]interface{}, hint),
	}
}

// Clear is removes all of the mappings from this map.
func (m *MutexMap) Clear() {
	m.Lock()
	m.items = make(map[string]interface{}, m.hint)
	m.Unlock()
}

// Clone is returns a shallow copy of this map instance.
// the keys and values themselves are not cloned.
func (m *MutexMap) Clone() *MutexMap {
	return m
}

// ContainsKey is returns true if this map contains a mapping for the specified key.
func (m *MutexMap) ContainsKey(key string) bool {
	m.RLock()
	_, ok := m.items[key]
	m.RUnlock()
	return ok
}

// ContainsValue is returns true if this map maps one or more keys to the specified value.
func (m *MutexMap) ContainsValue(item interface{}) bool {
	m.RLock()
	for k := range m.items {
		if m.items[k] == item {
			return true
		}
	}
	m.RUnlock()
	return false
}

// Equals is compares the specified object with this map for equality.
func (m *MutexMap) Equals(src map[string]interface{}) bool {
	m.RLock()
	defer m.RUnlock()
	if len(m.items) != len(src) {
		return false
	}
	for k := range m.items {
		if m.items[k] != src[k] {
			return false
		}
	}
	return true
}

// Get is returns the value to which the specified key is mapped, or nil
// if this map contains no mapping for the key.
func (m *MutexMap) Get(key string) interface{} {
	m.RLock()
	item := m.items[key]
	m.RUnlock()
	return item
}

// GetAll is returns the this map.
func (m *MutexMap) GetAll() map[string]interface{} {
	m.RLock()
	items := make(map[string]interface{}, len(m.items))
	for k, v := range m.items {
		items[k] = v
	}
	m.RUnlock()
	return items
}

// IsEmpty is returns true if this map contains no key-value mappings.
func (m *MutexMap) IsEmpty() bool {
	return m.Len() == 0
}

// KeySet is returns the keys contained in this map.
func (m *MutexMap) KeySet() []string {
	m.RLock()
	items := make([]string, 0, len(m.items))
	for k := range m.items {
		items = append(items, k)
	}
	m.RUnlock()
	return items
}

// Put is associates the specified value with the specified key in this map.
func (m *MutexMap) Put(key string, items interface{}) {
	m.Lock()
	m.items[key] = items
	m.Unlock()
}

// PutAll is copies all of the mappings from the specified map to this map.
func (m *MutexMap) PutAll(src map[string]interface{}) {
	m.Lock()
	m.items = src
	m.Unlock()
}

// Delete is deletes the mapping for a key from this map if it is present.
func (m *MutexMap) Delete(key string) {
	m.Lock()
	delete(m.items, key)
	m.Unlock()
}

// Len is Returns the number of key-value mappings in this map.
func (m *MutexMap) Len() int {
	m.RLock()
	l := len(m.items)
	m.RUnlock()
	return l
}

// Values is returns a list of the values contained in this map.
func (m *MutexMap) Values() []interface{} {
	m.RLock()
	items := make([]interface{}, 0, len(m.items))
	for _, v := range m.items {
		items = append(items, v)
	}
	m.RUnlock()
	return items
}
