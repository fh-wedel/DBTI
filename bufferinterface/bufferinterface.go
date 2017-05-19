// Package fileinterface provides access to the underlying operating system files.
// Files are organized in blocks.
package bufferinterface

import (
	"errors"
	"github.com/fh-wedel/DBTI/fileinterface"
	"math/rand"
)

const PageSize = fileinterface.Blocksize

type Page [PageSize]byte

const Buffersize = 16

var Buffer [Buffersize]Page

func searchBuffer(pageNo int) int {
	for i := 0; i < Buffersize; i++ {
		b := &Buffer[i]
		var no int = (int(b[0])*0x01000000 + int(b[1])*0x00010000 + int(b[2])*0x00000100 + int(b[3])*0x00000001)
		if no == pageNo {
			return i
		}
	}
	return -1
}

func allocateNewFrame(pageNo int) int {
	no := rand.Intn(Buffersize)
	b := &Buffer[no]
	// Record pageno
	b[0] = byte((pageNo & 0xFF000000) >> 24)
	b[1] = byte((pageNo & 0x00FF0000) >> 16)
	b[2] = byte((pageNo & 0x0000FF00) >> 8)
	b[3] = byte((pageNo & 0x000000FF))
	return no
}

// Request page with number pageNo. Returns pointer to page data in system buffer (and err is nil)
// If unsuccessful, return nil and an error value describing the error.
func Request(pageNo int) (*Page, error) {
	var nr int = searchBuffer(pageNo)
	if nr == -1 {
		nr = allocateNewFrame(pageNo)
		if nr == -1 {
			return nil, errors.New("Buffer exhausted!")
		}
	}

	return &Buffer[nr], nil
}

// Mark the page pageNo as pinned. It's pointer will stay valid and the page
// is nerver removed from the system bufffer.
// If unsuccessful, return an error value describing the error.
func Fix(pageNo int) error {
	return errors.New("not implemented")
}

// Mark the page pageNo as no longer pinned. The page might be subsequently
// removed from the system buffer.
// If unsuccessful, return an error value describing the error.
func UnFix(pageNo int) error {
	return errors.New("not implemented")
}

// Mark the page pageNo as modified. If the page ist later removed form the
// system buffer, it has to be written to mass storage.
// If unsuccessful, return an error value describing the error.
func Update(blockNo int) error {
	return errors.New("not implemented")
}

// Force an immediate write of this page to mass storage.
// The page address stays valid.
// If unsuccessful, return an error value describing the error.
func Write(pageNo int) error {
	return errors.New("not implemented")
}

/* ... */
