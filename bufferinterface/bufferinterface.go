// Package fileinterface provides access to the underlying operating system files.
// Files are organized in blocks.
package bufferinterface

import (
	"errors"
	"github.com/fh-wedel/DBTI/fileinterface"
	// "log"
)

const PageSize = fileinterface.Blocksize

type Page [PageSize]byte

// Request page with number pageNo. Returns pointer to page data in system buffer (and err is nil)
// If unsuccessful, return nil and an error value describing the error.
func Request(pageNo int) (*Page, error) {
	// log.Printf("Request %d", pageNo)
	var no int = searchBuffer(pageNo)
	if no == -1 { // page not found in buffer
		no = allocateNewPage(pageNo)
		if no == -1 { // cannot allocate fresh page: replace one
			if togo, err := s.pageToReplace(); err != nil {
				return nil, err
			} else if err = freePage(togo); err != nil {
				return nil, errors.New("Buffer exhausted!")
			} else {
				no = togo
			}
		}
	} else { // page found in buffer
		// log.Printf("Found page %v in buffer at %v", pageNo, no)
		reference(no)
		return &s.buffer[no], nil
	}

	// page in buffer
	page := &s.buffer[no]

	// calculate block no, direct block adressing used here
	blockNo := pageNo

	// read the block
	block, err := fileinterface.Read(s.currentFile, blockNo)
	if err != nil {
		return nil, err
	}

	// place it into buffer
	block2page(block, page)

	// set page header
	setFixed(page, false)
	setModified(page, false)
	setBlockNo(page, pageNo)

	reference(no)

	return page, nil
}

// Mark the page pageNo as pinned. It's pointer will stay valid and the page
// is never removed from the system bufffer.
// If unsuccessful, return an error value describing the error.
func Fix(pageNo int) error {
	if nr := searchBuffer(pageNo); nr < 0 {
		return errors.New("no such page - request page first")
	} else {
		p := &s.buffer[nr]
		setFixed(p, true)
		return nil
	}
}

// Mark the page pageNo as no longer pinned. The page might be subsequently
// removed from the system buffer.
// If unsuccessful, return an error value describing the error.
func UnFix(pageNo int) error {
	if nr := searchBuffer(pageNo); nr < 0 {
		return errors.New("no such page - request page first")
	} else {
		p := &s.buffer[nr]
		if isFixed(p) {
			setFixed(p, false)
			return nil
		}
		return errors.New("page has not been fixed before")
	}
}

// Mark the page pageNo as modified. If the page ist later removed form the
// system buffer, it has to be written to mass storage.
// If unsuccessful, return an error value describing the error.
func Update(pageNo int) error {
	if nr := searchBuffer(pageNo); nr < 0 {
		return errors.New("no such page - request page first")
	} else {
		p := &s.buffer[nr]
		setModified(p, true)
		return nil
	}
}

// Force an immediate write of this page to mass storage.
// The page address stays valid.
// If unsuccessful, return an error value describing the error.
func Write(pageNo int) error {
	if nr := searchBuffer(pageNo); nr < 0 {
		return errors.New("no such page - request page first")
	} else {
		p := &s.buffer[nr]
		if isModified(p) {
			if err := writePage(p); err != nil {
				return err
			}
		}
		setModified(p, false)
		return nil
	}
}

// Set the ReplacementStrategy
func SetReplacementStrategy(strategy replacementStrategy) error {
	s.pageToReplace = strategy
	return nil
}
