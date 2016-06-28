package superhash

import (
	"testing"
)

func TestSuperHashSetGet(t *testing.T) {
	hashmap := New()
	k1, k2, k3, value := 1, true, 3, 4
	hashmap.Set(k1, k2, k3, value)
	if hashmap.Get(k1, k2, k3) != value {
		t.Error("invalid value for keys")
	}
}

func TestSuperHashSetGetExistingPath(t *testing.T) {
	hashmap := New()
	k1, k2, k3, v, v2 := 1, true, 3, 4, 5
	hashmap.Set(k1, k2, k3, v)
	hashmap.Set(k1, k2, v2)
	if hashmap.Get(k1, k2) != v2 {
		t.Error("invalid value for keys")
	}
}

func TestSuperHashSetWithNoValue(t *testing.T) {
	hashmap := New()
	if hashmap.Set("test") != false {
		t.Error("Set with no value should return false")
	}
}

func TestSuperHashGetWithNoKeys(t *testing.T) {
	hashmap := New()
	k1, k2, k3, value := 1, true, 3, 4
	hashmap.Set(k1, k2, k3, value)
	hashmap.Get()
}

func TestSuperHashDelete(t *testing.T) {
	hashmap := New()
	k1, k2, k3, value := 1, true, 3, 4
	hashmap.Set(k1, k2, k3, value)
	hashmap.Delete(k1, k2, k3)
	if hashmap.Get(k1, k2, k3) != nil {
		t.Error("deleted keys still accessible")
	}
}

func TestSuperHashDeleteWithNoKeys(t *testing.T) {
	hashmap := New()
	k1, k2, k3, value := 1, true, 3, 4
	hashmap.Set(k1, k2, k3, value)
	hashmap.Delete()
}

func TestSuperHashSetGetConcurrent(t *testing.T) {
	hashmap := New()
	k1, k2, k3 := 1, true, 3
	count := 0
	for count < 1000 {
		go func() {
			count++
			hashmap.Set(k1, k2, k3, count)
		}()
		go func() {
			count++
			hashmap.Set(k1, k2, k3, count)
		}()
		go func() {
			count++
			hashmap.Set(k1, k2, k3, count)
		}()
		go func() {
			count++
			hashmap.Set(k1, k2, k3, count)
		}()
	}
}