package btree

import (
	// "errors"
	"fmt"
	"testing"
)

func check(t *testing.T, bt *Btree, key string, expected int) {
	i, found := bt.Find(key)
	if !found {
		t.Errorf("Value %v not found although inserted", key)
	} else if i != expected {
		t.Errorf("Expected %v but found %v", expected, i)
	}
}

func checkMissing(t *testing.T, bt *Btree, key string) {
	_, found := bt.Find(key)
	if found {
		t.Errorf("Value %v found although not present", key)
	}
}
func TestNew(t *testing.T) {
	NewBtree("a", 1, "b", 2)
}

func TestPrint1(t *testing.T) {
	bt := NewBtree("a", 1, "b", 2)

	fmt.Println(bt)
}

func TestPrint2(t *testing.T) {
	b0 := Btree{z0: nil, entries: [2 * k]entry{entry{"a", 1, nil}, entry{"b", 2, nil}}}
	b1 := Btree{z0: nil, entries: [2 * k]entry{entry{"d", 4, nil}, entry{"e", 5, nil}}}
	b2 := Btree{z0: nil, entries: [2 * k]entry{entry{"g", 7, nil}, entry{"h", 8, nil}}}

	bt := &Btree{z0: &b0, entries: [2 * k]entry{entry{"c", 3, &b1}, entry{"f", 6, &b2}}}

	fmt.Println(bt)
}

func TestLength1(t *testing.T) {

}

func TestLength2(t *testing.T) {

}

func TestFind1(t *testing.T) {
	b0 := Btree{z0: nil, entries: [2 * k]entry{entry{"a", 1, nil}, entry{"b", 2, nil}}}
	b1 := Btree{z0: nil, entries: [2 * k]entry{entry{"d", 4, nil}, entry{"e", 5, nil}}}
	b2 := Btree{z0: nil, entries: [2 * k]entry{entry{"g", 7, nil}, entry{"h", 8, nil}}}

	bt := &Btree{z0: &b0, entries: [2 * k]entry{entry{"c", 3, &b1}, entry{"f", 6, &b2}}}

	check(t, bt, "a", 1)
	check(t, bt, "b", 2)
	check(t, bt, "c", 3)
	check(t, bt, "d", 4)
	check(t, bt, "e", 5)
	check(t, bt, "f", 6)
	check(t, bt, "g", 7)
	check(t, bt, "h", 8)

	checkMissing(t, bt, "x")

}

func TestFind2(t *testing.T) {

	b0 := Btree{z0: nil, entries: [2 * k]entry{entry{"a", 1, nil}, entry{"b", 2, nil},
		entry{"b2", 1, nil}, entry{"b3", 2, nil},
	}}
	b1 := Btree{z0: nil, entries: [2 * k]entry{entry{"d", 4, nil}, entry{"e", 5, nil}}}
	b2 := Btree{z0: nil, entries: [2 * k]entry{entry{"g", 7, nil}, entry{"h", 8, nil}}}

	bt := &Btree{z0: &b0, entries: [2 * k]entry{entry{"c", 3, &b1}, entry{"f", 6, &b2}}}

	check(t, bt, "a", 1)
	check(t, bt, "b", 2)
	check(t, bt, "c", 3)
	check(t, bt, "d", 4)
	check(t, bt, "e", 5)
	check(t, bt, "f", 6)
	check(t, bt, "g", 7)
	check(t, bt, "h", 8)

	check(t, bt, "b2", 1)
	check(t, bt, "b3", 2)

	checkMissing(t, bt, "x")
}

func TestFind3(t *testing.T) {
}

func TestInsertExists(t *testing.T) {
	bt := NewBtree("a", 1, "b", 2)

	x := bt.Insert("a", 3)

	if x != nil {
		t.Error("Got Tree expecting nil")
	}

}

func TestInsert1(t *testing.T) {
	bt := NewBtree("a", 1, "b", 2)
	checkMissing(t, bt, "c")

	bt.Insert("c", 3)
	check(t, bt, "c", 3)

	// fmt.Println(bt)
}

func TestInsert2(t *testing.T) {
	bt := NewBtree("a", 1, "b", 2)
	checkMissing(t, bt, "c")
	checkMissing(t, bt, "d")

	bt.Insert("c", 3)
	check(t, bt, "c", 3)
	bt.Insert("d", 4)

	check(t, bt, "a", 1)
	check(t, bt, "b", 2)
	check(t, bt, "c", 3)
	check(t, bt, "d", 4)
	// fmt.Println(bt)
}

func TestInsert2a(t *testing.T) { // check entries not inserted in order
	bt := NewBtree("a", 1, "b", 2)
	checkMissing(t, bt, "c")
	checkMissing(t, bt, "d")

	bt.Insert("d", 4)
	check(t, bt, "d", 4)
	bt.Insert("c", 3)

	check(t, bt, "a", 1)
	check(t, bt, "b", 2)
	check(t, bt, "c", 3)
	check(t, bt, "d", 4)
	// fmt.Println(bt)
}

func TestInsert3(t *testing.T) {
	bt := NewBtree("a", 1, "b", 2)
	bt.Insert("c", 3)
	bt.Insert("d", 4)
	checkMissing(t, bt, "e")

	bt.Insert("e", 5) // node overflow, rebalance necessary

	check(t, bt, "a", 1)
	check(t, bt, "b", 2)
	check(t, bt, "c", 3)
	check(t, bt, "d", 4)
	check(t, bt, "e", 5)
}

func TestDelete1(t *testing.T) {
}

func TestDelete2(t *testing.T) {
}

func TestDelete3(t *testing.T) {
}

/* ... */
