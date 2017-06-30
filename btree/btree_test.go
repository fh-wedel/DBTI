package btree

import (
	// "errors"
	"fmt"
	"testing"
)

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

	i, found := bt.Find("a")
	if !found {
		t.Error("Value a not found although present")
	} else if i != 1 {
		t.Errorf("Expected 1 but found %v", i)
	}

	i, found = bt.Find("b")
	if !found {
		t.Error("Value b not found although present")
	} else if i != 2 {
		t.Errorf("Expected 2 but found %v", i)
	}

	i, found = bt.Find("c")
	if !found {
		t.Error("Value c not found although present")
	} else if i != 3 {
		t.Errorf("Expected 3 but found %v", i)
	}

	i, found = bt.Find("d")
	if !found {
		t.Error("Value d not found although present")
	} else if i != 4 {
		t.Errorf("Expected 4 but found %v", i)
	}

	i, found = bt.Find("e")
	if !found {
		t.Error("Value e not found although present")
	} else if i != 5 {
		t.Errorf("Expected 5 but found %v", i)
	}

	i, found = bt.Find("f")
	if !found {
		t.Error("Value f not found although present")
	} else if i != 6 {
		t.Errorf("Expected 6 but found %v", i)
	}

	i, found = bt.Find("g")
	if !found {
		t.Error("Value g not found although present")
	} else if i != 7 {
		t.Errorf("Expected 7 but found %v", i)
	}

	i, found = bt.Find("h")
	if !found {
		t.Error("Value h not found although present")
	} else if i != 8 {
		t.Errorf("Expected 8 but found %v", i)
	}

	i, found = bt.Find("x")
	if found {
		t.Error("Value x found although not present")
	}

}

func TestFind2(t *testing.T) {

	b0 := Btree{z0: nil, entries: [2 * k]entry{entry{"a", 1, nil}, entry{"b", 2, nil},
		entry{"b2", 1, nil}, entry{"b3", 2, nil},
	}}
	b1 := Btree{z0: nil, entries: [2 * k]entry{entry{"d", 4, nil}, entry{"e", 5, nil}}}
	b2 := Btree{z0: nil, entries: [2 * k]entry{entry{"g", 7, nil}, entry{"h", 8, nil}}}

	bt := &Btree{z0: &b0, entries: [2 * k]entry{entry{"c", 3, &b1}, entry{"f", 6, &b2}}}

	i, found := bt.Find("a")
	if !found {
		t.Error("Value a not found although present")
	} else if i != 1 {
		t.Errorf("Expected 1 but found %v", i)
	}

	i, found = bt.Find("b")
	if !found {
		t.Error("Value b not found although present")
	} else if i != 2 {
		t.Errorf("Expected 2 but found %v", i)
	}

	i, found = bt.Find("c")
	if !found {
		t.Error("Value c not found although present")
	} else if i != 3 {
		t.Errorf("Expected 3 but found %v", i)
	}

	i, found = bt.Find("d")
	if !found {
		t.Error("Value d not found although present")
	} else if i != 4 {
		t.Errorf("Expected 4 but found %v", i)
	}

	i, found = bt.Find("e")
	if !found {
		t.Error("Value e not found although present")
	} else if i != 5 {
		t.Errorf("Expected 5 but found %v", i)
	}

	i, found = bt.Find("f")
	if !found {
		t.Error("Value f not found although present")
	} else if i != 6 {
		t.Errorf("Expected 6 but found %v", i)
	}

	i, found = bt.Find("g")
	if !found {
		t.Error("Value g not found although present")
	} else if i != 7 {
		t.Errorf("Expected 7 but found %v", i)
	}

	i, found = bt.Find("h")
	if !found {
		t.Error("Value h not found although present")
	} else if i != 8 {
		t.Errorf("Expected 8 but found %v", i)
	}

	i, found = bt.Find("b2")
	if !found {
		t.Error("Value h not found although present")
	} else if i != 1 {
		t.Errorf("Expected 1 but found %v", i)
	}

	i, found = bt.Find("b3")
	if !found {
		t.Error("Value h not found although present")
	} else if i != 2 {
		t.Errorf("Expected 2 but found %v", i)
	}

	i, found = bt.Find("x")
	if found {
		t.Error("Value x found although not present")
	}

}

func TestFind3(t *testing.T) {
}

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
