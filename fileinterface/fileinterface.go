// Package fileinterface provides access to the underlying operating system files.
// Files are organized in blocks.
package fileinterface

import (
	"errors"
)

const Blocksize = 4096

type FID int // File ID

type Block [Blocksize]byte

// Creates a new file with a given name. Return the FID and nil if succesful.
// The file is then open for reading and writing.
// If unsuccessful, return any FID and an error value describing the error.
func Create(name string) (FID, error) {
	return 0, errors.New("not implemented")
}

// Deletes a file with a given name. Return nil if succesful.
// If unsuccessful, return an error value describing the error.
func Delete(name string) error {
	return errors.New("not implemented")
}

// Opens a file with a given name for reading and writing. Return nil if succesful.
// If unsuccessful, return  any FID and an error value describing the error.
// Possible errors include FileNotFoundError oder FileAlreadyOpenError
func Open(name string) (FID, error) {
	return 0, errors.New("not implemented")
}

// Length calculates the number of blocks available in the file given by fileNo. Return nil if succesful.
// If unsuccessful, return  any FID and an error value describing the error.
// Possible errors include FileNotOpenError
func Length(fileNo FID) (int, error) {
	return 0, errors.New("not implemented")
}

// Reads the block number blockNo from the file fileNo. Counting starts at 0.
// Return a pointer to the block and nil if succesful.
// If unsuccessful, return nil and an error value describing the error.
// Possible errors include FileNotOpenError
func Read(fileNo FID, blockNo int) (*Block, error) {
	return &Block{}, errors.New("not implemented")
}

// Writes the block given by the pointer block to the block number blockNo in
// the file fileNo. Counting starts at 0.
// Return nil if succesful.
// If unsuccessful, return an error value describing the error.
// Possible errors include FileNotOpen or WriteError
func Write(fileNo FID, blockNo int, block *Block) error { //  FileNotOpenException, IOException;
	return errors.New("not implemented")
}

// Close the file given by fileNo. Return nil if succesful.
// If unsuccessful, return an error value describing the error.
// Possible errors include FileNotOpenError
func Close(fileNo FID) error {
	return errors.New("not implemented")
}

// Return a list of possible file names
func ls() []string {
	return []string{}
}

/* ... */
