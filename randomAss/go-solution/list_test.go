package main

import "testing"

func TestAddFront(t *testing.T) {
	list := NewList[int]()

	list.AddFront(10)
	list.AddFront(20)
	list.AddFront(30)

	if list.GetSize() != 3 {
		t.Errorf("expected size 3, got %d", list.GetSize())
	}

	if list.GetHead().GetData() != 30 {
		t.Errorf("expected head 30, got %d", list.GetHead().GetData())
	}

	if list.GetTail().GetData() != 10 {
		t.Errorf("expected tail 10, got %d", list.GetTail().GetData())
	}
}

func TestAddBack(t *testing.T) {
	list := NewList[string]()

	list.AddBack("A")
	list.AddBack("B")
	list.AddBack("C")

	if list.GetSize() != 3 {
		t.Errorf("expected size 3, got %d", list.GetSize())
	}

	if list.GetHead().GetData() != "A" {
		t.Errorf("expected head A, got %s", list.GetHead().GetData())
	}

	if list.GetTail().GetData() != "C" {
		t.Errorf("expected tail C, got %s", list.GetTail().GetData())
	}
}

func TestAddFrontAndBack(t *testing.T) {
	list := NewList[int]()

	list.AddFront(1)
	list.AddBack(2)
	list.AddFront(3)

	if list.GetSize() != 3 {
		t.Errorf("expected size 3, got %d", list.GetSize())
	}

	if list.GetHead().GetData() != 3 {
		t.Errorf("expected head 3, got %d", list.GetHead().GetData())
	}

	if list.GetTail().GetData() != 2 {
		t.Errorf("expected tail 2, got %d", list.GetTail().GetData())
	}
}

func TestEmptyList(t *testing.T) {
	list := NewList[float64]()

	if list.GetHead() != nil {
		t.Error("expected head nil for empty list")
	}

	if list.GetTail() != nil {
		t.Error("expected tail nil for empty list")
	}

	if list.GetSize() != 0 {
		t.Errorf("expected size 0, got %d", list.GetSize())
	}
}
