package bufferinterface

import (
	"errors"
	"fmt"
	"github.com/fh-wedel/DBTI/fileinterface"
	// "log"
	"strconv"
)

// Pages are structured like this:
// { modified flag 1 byte: bit 0 modifed, bit 1 fixed | blockNo in 4 bytes (big endian) -1 if free | netto data ... }
func getBlockNo(p *Page) int {
	return int(p[1])<<24 + int(p[2])<<16 + int(p[3])<<8 + int(p[4])
}

// Set the block number of page
func setBlockNo(p *Page, blockNo int) {
	p[1] = byte((blockNo & 0xFF000000) >> 24)
	p[2] = byte((blockNo & 0x00FF0000) >> 16)
	p[3] = byte((blockNo & 0x0000FF00) >> 8)
	p[4] = byte((blockNo & 0x000000FF))
}

// Look for page with given blockNo within the buffer, return buffer index
// if succesful. Return -1 if page is not found.
func searchBuffer(blockNo int) int {
	for i := 0; i < buffersize; i++ {
		if getBlockNo(&s.buffer[i]) == blockNo {
			return i
		}
	}
	return -1
}

// Check if the given page is allocated to any block.
func isAllocated(p *Page) bool {
	return getBlockNo(p) != 0xFFFFFFFF
}

// Check if the given page has been modified
func isModified(p *Page) bool {
	return p[0]&1 != 0
}

// Set the modified state of the given page.
func setModified(p *Page, modified bool) {
	if modified {
		p[0] |= 1
	} else {
		p[0] &= (^1 & 0xFF)
	}
}

// Check if the given page is fixed (pinned)
func isFixed(p *Page) bool {
	return p[0]&2 != 0
}

// Set the fixed state of the given page.
func setFixed(p *Page, fixed bool) {
	if fixed {
		p[0] |= 2
	} else {
		p[0] &= (^2 & 0xFF)
	}
}

// Choose a new empty page from the buffer
func allocateNewPage(pageNo int) int {
	// log.Printf("allocateNewPage %d", pageNo)
	for i := 0; i < buffersize; i++ {
		p := &s.buffer[i]
		if !isAllocated(p) {
			// log.Printf("allocateNewPage page %d is free", i)
			return i
		}
	}
	// log.Printf("allocateNewPage no free page")
	return -1
}

// Make page with ginen pageNo available
func freePage(no int) error {
	// log.Printf("freePage %d", no)
	// see if frame is allocate
	page := &s.buffer[no]
	if !isAllocated(page) {
		return errors.New("Page " + strconv.Itoa(no) + " is not allocated!")
	}
	if isFixed(page) {
		return errors.New("Page is fixed.")
	}
	// log.Printf(" page[0]=%d", page[0])

	// see if page is modified
	if isModified(page) { // write old content
		// log.Printf("modified page %d", no)
		if err := writePage(page); err != nil {
			return err
		}
		setModified(page, false)
	}
	setBlockNo(page, -1)
	return nil
}

// copy block to page
func block2page(block *fileinterface.Block, page *Page) {
	for i := 0; i < PageSize; i++ {
		page[i] = block[i]
	}
}

// copy page to block
func page2block(page *Page, block *fileinterface.Block) {
	for i := 0; i < PageSize; i++ {
		block[i] = page[i]
	}
}

// write page to file
func writePage(page *Page) error {
	blockNo := getBlockNo(page)
	var block fileinterface.Block
	page2block(page, &block)
	err := fileinterface.Write(s.currentFile, blockNo, &block)
	return err
}

// free all pages writing all modified pages
func flush() error {
	for i, _ := range s.buffer {
		if err := freePage(i); err != nil {
			return err
		}
	}
	return nil
}

// initialize buffer
func emptyBuffer() {
	for i := 0; i < buffersize; i++ {
		p := &s.buffer[i]
		setFixed(p, false)
		setModified(p, false)
		setBlockNo(p, -1)
	}
}

// display the buffer structure
func showBuffer() {
	for i, p := range s.buffer {
		fmt.Printf("%d: %d ", i, getBlockNo(&p))
		if isAllocated(&p) {
			fmt.Printf(" allocated")
		} else {
			fmt.Printf(" free")
		}
		if isModified(&p) {
			fmt.Printf(" modified")
		} else {
			fmt.Printf(" unmodified")
		}
		if isFixed(&p) {
			fmt.Printf(" fixed")
		} else {
			fmt.Printf(" unfixed")
		}
		fmt.Printf("\n")
	}
}

// record that a page is referenced
func reference(pageNo int) {
	// log.Printf("references %v", pageNo)

	// for lruReplacementStrategy
	// if previously referenced, remove it
	for i, el := range s.references {
		if el == pageNo {
			s.references = append(s.references[:i], s.references[i+1:]...)
			break
		}
	}
	s.references = append(s.references, pageNo)
	// log.Printf("references %v", s.references)

	// for clockReplacementStategy
	s.counters[s.clock] = 1
}
