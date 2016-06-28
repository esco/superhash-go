package superhash

import (
	"sync"
)

type Node struct {
	key      interface{}
	value    interface{}
	children map[interface{}]*Node
}

func newNode() (n *Node) {
	n = &Node{}
	n.children = make(map[interface{}]*Node)
	return
}

type SuperHash struct {
	root *Node
	mutex *sync.Mutex
}

func New() (s *SuperHash) {
	s = &SuperHash{
		root: newNode(),
		mutex: &sync.Mutex{},
	}
	return
}

func (s *SuperHash) Set(kv ...interface{}) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(kv) < 2 {
		return false
	}
	keys := kv[:len(kv)-1]
	value := kv[len(kv)-1:][0]
	cn := s.root
	for _, key := range keys {
		if n, ok := cn.children[key]; ok {
			cn = n
		} else {
			n := newNode()
			n.key = key
			cn.children[key] = n
			cn = n
		}
	}
	cn.value = value
	return true
}

func (s *SuperHash) Get(keys ...interface{}) interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(keys) < 1 {
		return nil
	}
	cn := s.root
	for _, key := range keys {
		if n, ok := cn.children[key]; ok {
			cn = n
		} else {
			return nil
		}
	}
	return cn.value
}

func (s *SuperHash) Delete(keys ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(keys) < 1 {
		return
	}
	cn := s.root
	deletePath(cn, keys...)
}

func deletePath(parent *Node, keys ...interface{}) {
	if len(keys) == 0 {
		return
	}
	if n, ok := parent.children[keys[0]]; ok {
		deletePath(n, keys[1:]...)
	}
	delete(parent.children, keys[0])
}
