package fileinterface

import "testing"

func TestCreate(t *testing.T) {
	var f FID
	var err error
	if f, err = Open("dummyfile"); err != nil {
		t.Error(err)
	} else {
		if err = Close(f); err != nil {
			t.Error(err)
		}
	}
}

func TestDelete(t *testing.T) {
}

func TestOpen(t *testing.T) {
}

func TestRead(t *testing.T) {
}

func TestLength(t *testing.T) {
}

func TestWrite(t *testing.T) {
}

func TestLS(t *testing.T) {
}
