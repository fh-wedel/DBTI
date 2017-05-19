package bufferinterface

import (
	"testing"
	"unsafe"
)

func TestRequest(t *testing.T) {
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
}

func TestUnfix(t *testing.T) {
}

func TestUpdate(t *testing.T) {
}

func TestWrite(t *testing.T) {
}
