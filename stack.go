package datas

import (
	"sync"

	"fmt"
)

type Stack interface {
	// Push Item on stack.
	Push(t Item)

	// Pops Item from stack.
	Pop() (Item, error)

	// Returns length of the stack.
	Len() int

	// Returns the members of the stack as a slice.
	ToSlice() []Item

	// Returns a channel of elements that you can
	// range over.
	Iter() <-chan Item

	// Returns an Iterator object that you can use
	// to range over.
	Iterator() *Iterator

	// Provides a standard string representation of
	// the current state of stack.
	String() string
}

type stack struct {
	items []Item
	lock  *sync.RWMutex
}

// NewStack creates and returns a reference to an empty stack. Operations
// on the resulting stack are thread-safe
func NewStack(items ...Item) Stack {
	s := &stack{make([]Item, 0), &sync.RWMutex{}}
	for _, item := range items {
		s.Push(item)
	}
	return s
}

// NewStackFromSlice creates and returns reference to a stack from
// existing slice. Operations on the resulting stack are thread-safe
func NewStackFromSlice(items []Item) Stack {
	s := NewStack(items...)
	return s
}

func (s *stack) Push(i Item) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, i)
}

func (s *stack) Pop() (Item, error) {
	s.lock.Lock()
	if len(s.items) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	defer s.lock.Unlock()
	i := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return i, nil
}

func (s *stack) Len() int {
	return len(s.items)
}

func (s *stack) Iter() <-chan Item {
	ch := make(chan Item)
	go s.publishItems(ch)
	return ch
}

func (s *stack) Iterator() *Iterator {
	iterator, ch, stopCh := newIterator()
	go s.publishItemsWithStopCh(ch, stopCh)
	return iterator
}

func (s *stack) publishItems(ch chan Item) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	defer close(ch)
	for _, elem := range s.items {
		ch <- elem
	}
}

func (s *stack) publishItemsWithStopCh(ch chan<- Item, stopCh <-chan struct{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	defer close(ch)
L:
	for _, elem := range s.items {
		select {
		case <-stopCh:
			break L
		case ch <- elem:
		}

	}
}

func (s *stack) ToSlice() []Item {
	s.lock.RLock()
	items := make([]Item, len(s.items))
	defer s.lock.RUnlock()
	for _, elem := range s.items {
		items = append(items, elem)
	}

	return items
}

func (s *stack) String() string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	r := fmt.Sprintf("%v", s.items)
	return r
}
