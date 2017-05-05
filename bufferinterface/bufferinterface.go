// Package fileinterface provides access to the underlying operating system files.
// Files are organized in blocks.
package fileinterface

import (
	"errors"
	"github.com/fh-wedel/DBTI/fileinterface"
)

const PageSize = fileinterface.Blocksize

type Page [PageSize]byte

// Request page with number pageNo. Returns pointer to page data in system buffer (and err is nil)
// If unsuccessful, return nil and an error value describing the error.
func Request(PageNo int) (*Page, error) {
	return nil, errors.New("not implemented")
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
