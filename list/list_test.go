package list

import "testing"

func TestList_Insert(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 100; i++ {
		err := list.Insert(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	list.ForEach(func(node *Node[int]) {
		print(node.value, " ")
	})
}

func TestList_InsertHead(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 100; i++ {
		err := list.InsertHead(i)
		if err != nil {
			t.Fatal(err)
		}
	}
	list.ForEach(func(node *Node[int]) {
		print(node.value, " ")
	})
}

func TestList_Remove(t *testing.T) {
	list := NewList[int]()
	for i := 0; i < 100; i++ {
		err := list.InsertHead(i)
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < 50; i++ {
		err := list.Remove()
		if err != nil {
			t.Fatal(err)
		}
	}

	if err := list.ForEach(func(node *Node[int]) {
		print(node.value, " ")
	}); err != nil {
		t.Fatal(err)
	}
}
