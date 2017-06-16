package bufferinterface

// Types and state

import (
	"github.com/fh-wedel/DBTI/fileinterface"
)

const buffersize = 16

type replacementStrategy func() (int, error)

type state struct {
	buffer        [buffersize]Page
	pageToReplace replacementStrategy
	currentFile   fileinterface.FID

	// for lruReplacementStrategy
	references []int

	// for clockReplacementStrategy
	clock    int
	counters [buffersize]int
}

// Package state and intialization
var s state

func init() {
	var err error
	s.currentFile, err = fileinterface.Open("sample.database")
	if err != nil {
		panic(err)
	}
	// SetReplacementStrategy(randomReplacementStrategy)
	// SetReplacementStrategy(lruReplacementStrategy)
	SetReplacementStrategy(clockReplacementStrategy)
	emptyBuffer()
	// showBuffer()
}
