package btree

import (
	// "errors"
	// "os"
	"fmt"
)

const k = 2

type entry struct {
	s string
	d int
	z *Btree
}

type Btree struct {
	z0      *Btree
	entries [2 * k]entry
}

func (n *Btree) length() int {
	for i := 0; i < 2*k; i++ {
		if n.entries[i].s == "" {
			return i
		}
	}
	return 2 * k
}

func (n *Btree) exhausted() bool {
	return n.length() >= 2*k
}

func NewBtree(s1 string, d1 int, s2 string, d2 int) *Btree {
	if s2 < s1 {
		s1, s2 = s2, s1
		d1, d2 = d2, d1
	}
	e1 := entry{s1, d1, nil}
	e2 := entry{s2, d2, nil}
	return &Btree{z0: nil, entries: [2 * k]entry{e1, e2}}
}

// search looks for a given key in a node
// return pointer to entry or nil if not found.
// Do not follow pointers in node
func (n *Btree) searchInNode(s string) *entry {
	len := n.length()
	for i := 0; i < len; i++ { // searchentry{s, d, nil}
		if n.entries[i].s == s { // found
			return &n.entries[i]
		} else if s < n.entries[i].s { // beyond possible place
			return nil
		}
	} // not found
	return nil
}

// findInNodes looks for a given key in a node and its subnodes.
// return pointer to entry or nil if not found.
func (n *Btree) findInNodes(s string) (*Btree, *entry) {

	if n.isLeaf() {
		return n, n.searchInNode(s)
	}

	if s < n.entries[0].s { // find in first pointer
		return n.z0.findInNodes(s)
	} else { // find in entries
		len := n.length()
		for i := 0; i < len-1; i++ {
			if s == n.entries[i].s {
				return n, &n.entries[i]
			} else if s < n.entries[i+1].s {
				return n.entries[i].z.findInNodes(s)
			}
		}
		if s == n.entries[len-1].s {
			return n, &n.entries[len-1]
		} else {
			return n.entries[len-1].z.findInNodes(s)
		}
	}
}

// Find searches a btree for a given key
// it returns the value stored for the key
// and a true flag if present. If the key
// could not been found, return 0 and false.
func (b *Btree) Find(s string) (int, bool) {
	_, ent := b.findInNodes(s)
	if ent != nil {
		return ent.d, true
	} else {
		return 0, false
	}
}

// Split node that should accept entry e

func (n *Btree) splitNode(e entry) (*Btree, *Btree) {
	return nil, nil
}

func (n *Btree) insertInNode(s string, d int) {
	len := n.length()
	e := entry{s, d, nil}
	if len < 2*k { // still space: insert entry in order
		for k := len; k > 0; k-- {
			if n.entries[k-1].s < s { // done copying: k is the place to insert
				n.entries[k] = e
				return
			} else { // move entry
				n.entries[k] = n.entries[k-1]
			}
		} // all moved
		n.entries[0] = e
	} else {
		n.splitNode(e)
		// split node an insert
	}
}

func (b *Btree) Insert(s string, d int) *Btree {
	n1, ent := b.findInNodes(s)
	if ent != nil { // already present
		return nil
	} else {
		n1.insertInNode(s, d)
	}
	return b
}

func (b *Btree) Delete(s string) *Btree {
	return nil
}

func (n *Btree) isLeaf() bool {
	return n.z0 == nil
}

func (n *Btree) String() string {
	s := ""
	l := n.length()
	if n.isLeaf() {
		s = s + fmt.Sprintf("[ -")
		for i := 0; i < l; i++ {
			s = s + fmt.Sprintf(" | \"%v\", %v, - ", n.entries[i].s, n.entries[i].d)
		}
		s = s + fmt.Sprintf("]")
	} else { // no leaf
		s = s + fmt.Sprintf("[ %v", n.z0)
		for i := 0; i < l; i++ {
			s = s + fmt.Sprintf(" | \"%v\", %v, %v", n.entries[i].s, n.entries[i].d, n.entries[i].z)
		}
		s = s + fmt.Sprintf("]")
	}
	return s
}

func (b *Btree) Height() int {
	return 0
}
