package datas_test

import (
	"testing"

	"sync"

	"github.com/andrzejsliwa/datas"
)

func Test_NewStackFromSlice(t *testing.T) {
	slice := []datas.Item{1, 2, 3}
	s := datas.NewStackFromSlice(slice)
	if size := s.Len(); size != 3 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func Test_NewStack(t *testing.T) {
	s := datas.NewStack()
	if size := s.Len(); size != 0 {
		t.Errorf("wrong count, expected 0 and got %d", size)
	}
}

func TestStack_String(t *testing.T) {
	s := datas.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	r := s.String()
	if expected := "[1 2 3]"; r != expected {
		t.Errorf("wrong string representation: %v expected %v", r, expected)
	}
}

func TestStack_PushAndLen(t *testing.T) {
	s := datas.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if size := s.Len(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
}

func TestStack_PopSucceed(t *testing.T) {
	s := datas.NewStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)

	value, err := s.Pop()
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != 3 {
		t.Errorf("wrong value, expected 3 and got %d", value)
	}
	if size := s.Len(); size != 2 {
		t.Errorf("wrong length, expected 3 and got %d", size)
	}
}

func TestStack_PopFailed(t *testing.T) {
	s := datas.NewStack()
	value, err := s.Pop()
	if err == nil && value != nil {
		t.Error("expected empty stack error")
	}
	if len(err.Error()) == 0 {
		t.Error("missing error message")
	}
}

// always run tests with: go test -race
func TestStack_ToSliceDeadlock(t *testing.T) {
	var wg sync.WaitGroup
	s := datas.NewStack()
	workers := 10
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				s.Push(1)
				s.ToSlice()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// always run tests with: go test -race
func TestStack_PopDeadlock(t *testing.T) {
	var wg sync.WaitGroup
	s := datas.NewStack()
	workers := 10
	wg.Add(workers)
	for i := 1; i <= workers; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				s.Push(j)
				s.Pop()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
