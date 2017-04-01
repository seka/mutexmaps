package mutexmultimap

import (
	"reflect"
	"testing"
)

func TestMutexMultiMapNew(t *testing.T) {
	m := New(1)
	if !m.IsEmpty() {
		t.Fatal("length should be 0:", m.Len())
	}
}

func TestMutexMultiMapClear(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if l := m.Len(); l != 1 {
		t.Fatal("invalid length:", l)
	}
	m.Clear()
	if l := m.Len(); l != 0 {
		t.Fatal("invalid length:", l)
	}
}

func TestMutexMultiMapContainsEntry(t *testing.T) {
	m := New(1)
	if ok := m.ContainsEntry("a", "value"); ok {
		t.Fatal("ContainsEntry should be false:", ok)
	}
	m.Put("a", "value")
	if ok := m.ContainsEntry("a", "value"); !ok {
		t.Fatal("ContainsEntry should be true:", ok)
	}
}

func TestMutexMultiMapContainsKey(t *testing.T) {
	m := New(1)
	if ok := m.ContainsKey("a"); ok {
		t.Fatal("ContainsEntry should be false:", ok)
	}
	m.Put("a", "value")
	if ok := m.ContainsKey("a"); !ok {
		t.Fatal("ContainsEntry should be true:", ok)
	}
}

func TestMutexMultiMapContainsValue(t *testing.T) {
	m := New(1)
	if ok := m.ContainsValue("value"); ok {
		t.Fatal("ContainsValue should be false:", ok)
	}
	m.Put("a", "value")
	if ok := m.ContainsValue("value"); !ok {
		t.Fatal("ContainsEntry should be true:", ok)
	}
}

func TestMutexMultiMapEntries(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if l := m.Len(); l != 1 {
		t.Fatal("mutexmap length should be 1:", l)
	}
	values := m.Entries()
	expected := map[string][]interface{}{
		"a": {"value"},
	}
	if !reflect.DeepEqual(values, expected) {
		t.Fatal("invalid values:", values, expected)
	}
}

func TestMutexMultiMapEquals(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	items1 := map[string][]interface{}{}
	if m.Equals(items1) {
		t.Fatal("Equals should be false", m.Entries())
	}
	items2 := map[string][]interface{}{
		"a": {"value"},
	}
	if !m.Equals(items2) {
		t.Fatal("invalid values:", m.Entries())
	}
}

func TestMutexMultiMapPut(t *testing.T) {
	m := New(1)
	m.Put("a", "value1")
	m.Put("a", "value2")
	values := m.Get("a")
	expected := []interface{}{"value1", "value2"}
	if !reflect.DeepEqual(values, expected) {
		t.Fatal("invalid values:", values, expected)
	}
}

func TestMutexMultiMapPutAll(t *testing.T) {
	m := New(1)
	values := []interface{}{"value1", "value2"}
	m.PutAll("a", values)
	items := m.Get("a")
	if !reflect.DeepEqual(items, values) {
		t.Fatal("invalid values:", items, values)
	}
}

func TestMutexMultiMapDelete(t *testing.T) {
	m := New(1)
	m.Put("a", "value1")
	m.Put("a", "value2")
	if l := m.Len(); l != 1 {
		t.Fatal("mutexmap length should be 1:", l)
	}
	m.Delete("a", "value1")
	actual := m.Get("a")
	expected := []interface{}{"value2"}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("invalid values:", actual, expected)
	}
}

func TestMutexMultiMapDeleteAllValues(t *testing.T) {
	m := New(1)
	m.Put("a", "value")
	if l := m.Len(); l != 1 {
		t.Fatal("mutexmap length should be 1:", l)
	}
	m.Delete("a", "value")
	if l := m.Len(); l != 0 {
		t.Fatal("mutexMultiMap length should be 0:", l)
	}
	values := m.Get("a")
	expected := []interface{}{}
	if !reflect.DeepEqual(values, expected) {
		t.Fatal("invalid values1:", values, expected)
	}
}

func TestMutexMultiMapDeleteAll(t *testing.T) {
	m := New(1)
	m.Put("a", "value1")
	if l := m.Len(); l != 1 {
		t.Fatal("mutexmap length should be 1:", l)
	}
	m.DeleteAll("a")
	actual := m.Get("a")
	expected := []interface{}{}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("invalid values:", actual, expected)
	}
}

func TestMutexMultiMapReplaceValues(t *testing.T) {
	m := New(1)
	m.Put("a", "value1")
	actual1 := m.Get("a")
	expected1 := []interface{}{"value1"}
	if !reflect.DeepEqual(actual1, expected1) {
		t.Fatal("invalid values:", actual1, expected1)
	}
	m.ReplaceValues("a", "value2")
	actual2 := m.Get("a")
	expected2 := []interface{}{"value2"}
	if !reflect.DeepEqual(actual2, expected2) {
		t.Fatal("invalid values:", actual2, expected2)
	}
}
