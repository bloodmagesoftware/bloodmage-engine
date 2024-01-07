package utils_test

import (
	"testing"

	"github.com/bloodmagesoftware/bloodmage-engine/engine/utils"
)

func TestStack(t *testing.T) {
	t.Parallel()
	s := utils.NewStack[int]()

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Errorf("Stack should have 3 elements")
	}

	if v, b := s.Pop(); *v != 3 || !b {
		t.Errorf("Stack should return 3")
	}

	if s.Len() != 2 {
		t.Errorf("Stack should have 2 elements")
	}

	if v, b := s.Pop(); *v != 2 || !b {
		t.Errorf("Stack should return 2")
	}

	if s.Len() != 1 {
		t.Errorf("Stack should have 1 element")
	}

	if v, b := s.Pop(); *v != 1 || !b {
		t.Errorf("Stack should return 1")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}

	if v, b := s.Pop(); v != nil || b {
		t.Errorf("Stack should return nil and false")
	}

	if s.Len() != 0 {
		t.Errorf("Stack should be empty")
	}
}
