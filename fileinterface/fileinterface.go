// Package fileinterface provides access to the underlying operating system files.
// Files are organized in blocks.
package fileinterface

import (
	"errors"
	"os"
)

const Blocksize = 4096

const MAXOPENFILES = 16

type FID int // File ID

type Block [Blocksize]byte

var files [MAXOPENFILES]*os.File
var lastFid = -1

// Creates a new file with a given name. Return the FID and nil if succesful.
// The file is then open for reading and writing.
// If unsuccessful, return any FID and an error value describing the error.
func Create(name string) (FID, error) {
	file, err := os.Create(name)
	if err != nil {
		return 0, err
	} else {
		lastFid++
		files[lastFid] = file
		return FID(lastFid), nil
	}
}

// Deletes a file with a given name. Return nil if succesful.
// If unsuccessful, return an error value describing the error.
func Delete(name string) error {
	return os.Remove(name)
}

// Opens a file with a given name for reading and writing. Return nil if succesful.
// If unsuccessful, return  any FID and an error value describing the error.
// Possible errors include FileNotFoundError oder FileAlreadyOpenError
func Open(name string) (FID, error) {
	file, err := os.OpenFile(name, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	} else {
		lastFid++
		files[lastFid] = file
		return FID(lastFid), nil
	}
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
	var block Block
if file, err := checkOpenFile(fileNo); err != nil {
		return nil, err
	} else if _,err:=file.Seek(int64(blockNo*Blocksize),0); err != nil {
			return nil, err
	} else if _,err:=file.Read([]byte((block)[:])); err != nil {
		    return nil,err
	} else {
	  return &block, nil
    }
}

// Writes the block given by the pointer block to the block number blockNo in
// the file fileNo. Counting starts at 0.
// Return nil if succesful.
// If unsuccessful, return an error value describing the error.
// Possible errors include FileNotOpen or WriteError
func Write(fileNo FID, blockNo int, block *Block) error {
if file, err := checkOpenFile(fileNo); err != nil {
		return err
	} else if _,err:=file.Seek(int64(blockNo*Blocksize),0); err != nil {
			return err
	} else if _,err:=file.Write([]byte((*block)[:])); err != nil {
		    return err
	} else {
	  return nil
    }
}

// Close the file given by fileNo. Return nil if succesful.
// If unsuccessful, return an error value describing the error.
// Possible errors include FileNotOpenError
func Close(fileNo FID) error {
	if file, err := checkOpenFile(fileNo); err != nil {
		return err
	} else {
		return file.Close()
	}
}

func checkOpenFile(fileNo FID) (*os.File, error) {
	f := int(fileNo)
	if f >= MAXOPENFILES {
		return nil, errors.New("illegal FID")
	}
	file := files[f]
	if file == nil {
		return nil, errors.New("file not open")
	}
	return file, nil
}

// Return a list of possible file names
func ls() []string {
	return []string{}
}

/* ... */
