package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_ReturnsEmptyStack(t *testing.T) {
	s := New[int]()
	assert.Truef(t, s.IsEmpty(), "New stack is not empty, size got=%d expected=%d", len(s.items), 0)
}

func Test_Push_IncrementsSize(t *testing.T) {
	s := New[int]()
	s.Push(1)
	assert.Equalf(t, s.Size(), 1, "Push did not increment size, got=%d expected=%d", s.Size(), 1)
}

func Test_Pop_DecrementsSize(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Pop()
	assert.Equalf(t, s.Size(), 0, "Pop did not decrement size, got=%d expected=%d", s.Size(), 0)
}

func Test_Pop_ReturnsCorrectItem(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item, _ := s.Pop()
	assert.Equalf(t, item, 3, "Pop did not return correct item, got=%d expected=%d", item, 3)
}

func Test_Pop_ReturnsErrorOnEmpty(t *testing.T) {
	s := New[int]()
	_, err := s.Pop()
	assert.NotNil(t, err, "Pop on empty stack did not return an error")
}

func Test_Pop2_DecrementsSizeBy2(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Pop2()
	assert.Equalf(t, s.Size(), 0, "Pop2 did not decrement size by 2, got=%d expected=%d", s.Size(), 0)
}

func Test_Pop2_ReturnsCorrectItems(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	item1, item2, _ := s.Pop2()
	assert.Equalf(t, item1, 3, "Pop2 did not return correct first item, got=%d expected=%d", item1, 3)
	assert.Equalf(t, item2, 2, "Pop2 did not return correct second item, got=%d expected=%d", item2, 2)
}

func Test_Pop2_ReturnsErrorOnSizeLessThanTwo(t *testing.T) {
	s := New[int]()
	s.Push(1)
	_, _, err := s.Pop2()
	assert.NotNilf(t, err, "Pop2 on stack of size %d did not return an error", s.Size())
}

func Test_Peek_ReturnsTopItem(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	item := s.Peek()
	assert.Equalf(t, item, 2, "Peek did not return correct item, got=%d expected=%d", item, 2)
}

func Test_Items_ReturnsArrayOfItems(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	items := s.Items()
	assert.ElementsMatchf(t, items, []int{1, 2, 3},
		"Items did not return correct items in stack, got=%v expected=%v", items, []int{1, 2, 3})
}
