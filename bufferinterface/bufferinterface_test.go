package bufferinterface

import (
	"log"
	"testing"
	"unsafe"
)

func TestRequest(t *testing.T) {
	log.Println("**** TestRequest")

	for i := 0; i < buffersize; i++ {
		Request(i)
	}

	p1, err := Request(12)
	if err != nil {
		t.Error(err)
		return
	}
	if p1 == nil {
		t.Error("got nil")
		return
	}
	p2, err := Request(12)
	if err != nil {
		t.Error(err)
		return
	}
	if p2 == nil {
		t.Error("got nil")
		return
	}

	if p1 != p2 {
		t.Error("requested addresses do not match: ", unsafe.Pointer(p1), ", ", unsafe.Pointer(p2))
	}

}

func TestFix(t *testing.T) {
	log.Println("**** TestFix")
	if p, err := Request(9); err != nil {
		t.Error(err)
		return
	} else if err = Fix(9); err != nil {
		t.Error(err)
		return
	} else if !isFixed(p) {
		t.Error("Page is not fixed")
		return
	} else {

		// now fill the buffer with pages, fixed page should stay
		for i := 0; i < 2*buffersize; i++ {
			if _, err := Request(i); err != nil {
				t.Error(err)
				return
			}
		}

		p2, err := Request(9)
		if err != nil {
			t.Error(err)
			return
		}
		if p2 == nil {
			t.Error("got nil")
			return
		}

		if p != p2 {
			t.Error("requested addresses do not match: ", unsafe.Pointer(p), ", ", unsafe.Pointer(p2))
			return
		}
	}

}

func TestUnfix(t *testing.T) {
	log.Println("**** TestUnfix")
	if p, err := Request(9); err != nil {
		t.Error(err)
		return
	} else if err = Fix(9); err != nil {
		t.Error(err)
		return
	} else if !isFixed(p) {
		t.Error("Page is not fixed")
		return
	} else if err = UnFix(9); err != nil {
		t.Error(err)
		return
	} else if isFixed(p) {
		t.Error("Page is fixed")
		return
	} else {

		// now fill the buffer with pages, unfixed page should go
		for i := 0; i < 2*buffersize; i++ {
			if _, err := Request(i); err != nil {
				t.Error(err)
				return
			}
		}

		p2, err := Request(9)
		if err != nil {
			t.Error(err)
			return
		}
		if p2 == nil {
			t.Error("got nil")
			return
		}

		if p == p2 {
			t.Error("requested addresses still match: ", unsafe.Pointer(p), ", ", unsafe.Pointer(p2))
			return
		}
	}
}

func TestUpdate(t *testing.T) {
	log.Println("**** TestUpdate")

	if p, err := Request(9); err != nil {
		t.Error(err)
		return
	} else if err = Update(9); err != nil {
		t.Error(err)
		return
	} else if !isModified(p) {
		t.Error("Page is not marked as modified")
		return
	}
}

func TestWrite(t *testing.T) {
	log.Println("**** TestWrite")

	if p, err := Request(9); err != nil {
		t.Error(err)
		return
	} else if err = Update(9); err != nil {
		t.Error(err)
		return
	} else if !isModified(p) {
		t.Error("Page is not marked as modified")
		return
	} else if err = Write(9); err != nil {
		t.Error(err)
		return
	} else if isModified(p) {
		t.Error("Page is still marked as modified after writing")
	}
}

func TestBlockFill(t *testing.T) {
	log.Println("**** TestBlockFill")

	for i := 0; i < 95; i++ {
		p, err := Request(i)
		if err != nil {
			t.Error(err)
			return
		}
		if p == nil {
			t.Error("Got nil")
			return
		}
		// mt.Printf("i=%d, p=~v", i, p)
		for j := 5; j < PageSize; j++ {
			p[j] = byte(i)
		}
		p[5] = '*'
		Update(i)
	}
	if err := flush(); err != nil {
		t.Error(err)
		return
	}
}
