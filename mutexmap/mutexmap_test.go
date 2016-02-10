package mutexmap

import (
	"reflect"
	"testing"
)

func TestMutexMapNew(t *testing.T) {
	m := New(1)
	if !m.IsEmpty() {
		t.Fatal("mutexMap IsEmpty should be true")
	}
}

func TestMutexMapClear(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if m.Len() != 1 {
		t.Fatal("mutexmap length should be 1:", m.Len())
	}
	m.Clear()
	if !m.IsEmpty() {
		t.Fatal("mutexmap IsEmpty should be true")
	}
}

func TestMutexMapClone(t *testing.T) {
	oldMap := New(1)
	v1 := reflect.ValueOf(oldMap)
	newMap := oldMap.Clone()
	v2 := reflect.ValueOf(newMap)
	if v1.Pointer() != v2.Pointer() {
		t.Fatal("mutexmap Clone unexpected value:", v1.Pointer(), v2.Pointer())
	}
}

func TestMutexMapContainsKey(t *testing.T) {
	m := New(1)
	if m.ContainsKey("a") {
		t.Fatal("mutexmap ContainsKey should be false")
	}
	m.Put("a", "value")
	if !m.ContainsKey("a") {
		t.Fatal("mutexmap ContainsKey should be true")
	}
}

func TestMutexMapContainsValue(t *testing.T) {
	m := New(1)
	if m.ContainsValue("value") {
		t.Fatal("mutexmap ContainsValue should be false")
	}
	m.Put("a", "value")
	if !m.ContainsValue("value") {
		t.Fatal("mutexmap ContainsKey should be true")
	}
}

func TestMutexMapEquals(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	items := map[string]interface{}{
		"a": "value",
	}
	if !m.Equals(items) {
		t.Fatal("mutexmap Equals should be true")
	}
}

func TestMutexMapGetAll(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if m.Len() != 1 {
		t.Fatal("mutexMap length should be 1:", m.Len())
	}
	values := m.GetAll()
	expected := map[string]interface{}{
		"a": "value",
	}
	if !reflect.DeepEqual(values, expected) {
		t.Fatal("invalid values:", values)
	}
}

func TestMutexMapKeySet(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	actual := m.KeySet()
	expected := []string{"a"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("unexpected values:", actual)
	}
}

func TestMutexMapPut(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if m.Len() != 1 {
		t.Fatal("mutexMap length should be 1:", m.Len())
	}
	value := m.Get("a")
	expected := "value"
	if value != expected {
		t.Fatal("invalid value:", value)
	}
}

func TestMutexMapPutDuplicateValue(t *testing.T) {
	m := New(1)
	m.Put("a", "value1")
	m.Put("a", "value2")
	values := m.GetAll()
	expected := map[string]interface{}{
		"a": "value2",
	}
	if !reflect.DeepEqual(values, expected) {
		t.Fatal("invalid values:", values)
	}
}

func TestMutexMapDelete(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if m.Len() != 1 {
		t.Fatal("mutexMap length should be 1:", m.Len())
	}
	m.Delete("a")
	if m.Len() != 0 {
		t.Fatal("mutexMap length should be 0:", m.Len())
	}
	value := m.Get("a")
	if value != nil {
		t.Fatal("invalid value:", value)
	}
}

func TestMutexMapValues(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	actual := m.Values()
	expected := []interface{}{"value"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("unexpected values:", actual)
	}
}
