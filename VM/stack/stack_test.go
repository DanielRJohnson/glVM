package stack

import "testing"

func Test_New_ReturnsEmptyStack(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Errorf("New stack is not empty. Size Expected=%d, Got=%d", 0, len(s.items))
	}
}

func Test_Push_IncrementsSize(t *testing.T) {
	s := New[int]()
	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Push did not increment size. Size Expected=%d, Got=%d", 1, s.Size())
	}
}

func Test_Pop_DecrementsSize(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Pop()
	if s.Size() != 0 {
		t.Errorf("Pop did not decrement size. Size Expected=%d, Got=%d", 0, s.Size())
	}
}

func Test_Pop_ReturnsCorrectItem(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item, _ := s.Pop()
	if item != 3 {
		t.Errorf("Pop did not return correct item. Item Expected=%d, Got=%d", 3, item)
	}
}

func Test_Pop_ReturnsErrorOnEmpty(t *testing.T) {
	s := New[int]()
	_, err := s.Pop()
	if err == nil {
		t.Errorf("Popping an empty stack did not return an error")
	}
}

func Test_Pop2_DecrementsSizeBy2(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Pop2()
	if s.Size() != 0 {
		t.Errorf("Pop2 did not decrement size by 2. Size Expected=%d, Got=%d", 0, s.Size())
	}
}

func Test_Pop2_ReturnsCorrectItems(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item1, item2, _ := s.Pop2()
	if item1 != 3 || item2 != 2 {
		t.Errorf("Pop2 did not return correct items. Items Expected=(%d, %d), Got=(%d, %d)", 3, 2, item1, item2)
	}
}

func Test_Pop2_ReturnsErrorOnSizeLessThanTwo(t *testing.T) {
	s := New[int]()
	s.Push(1)
	_, _, err := s.Pop2()
	if err == nil {
		t.Errorf("Pop-2-ing a stack of size %d did not return an error", s.Size())
	}
}
