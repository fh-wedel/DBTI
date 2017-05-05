package fileinterface

import (
  "testing"
  "errors"
  )

func TestCreate(t *testing.T) {
	var f FID
	var err error
	if f, err = Create("dummyfile"); err != nil {
		t.Error(err)
	} else {
		if err = Close(f); err != nil {
			t.Error(err)
		}
	}
}

func TestDelete(t *testing.T) {
	err := Delete("dummyfile")
	if err != nil {
		t.Error(err)
	}
}

func TestOpen(t *testing.T) {
	if f, err := Open("sample.database"); err != nil {
		t.Error(err)
	} else if err = Close(f); err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	var block *Block
	if f, err := Open("sample.database"); err != nil {
		t.Error(err)
	} else if block, err = Read(f, 5); err != nil {
		t.Error(err)
	} else if err = Close(f); err != nil {
		t.Error(err)
	}
	for i:=0; i<Blocksize; i++ {
		if block[i] != byte(i & 0xFF) {
			t.Error(errors.New("block has unexpected value"))
			break
		}
	}
}

func TestLength(t *testing.T) {
}

func TestWrite(t *testing.T) {
	var block Block
	for i:=0; i<Blocksize; i++ {
		block[i]= byte(i & 0xFF)
	}
	if f, err := Open("sample.database"); err != nil {
		t.Error(err)
	} else if err := Write(f, 5, &block); err != nil {
		t.Error(err)
	} else if err = Close(f); err != nil {
		t.Error(err)
	}
}

func TestLS(t *testing.T) {
}
