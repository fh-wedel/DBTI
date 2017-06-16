package bufferinterface

import (
	"errors"
	// "log"
	"math/rand"
)

// Choose a page to replace random
func randomReplacementStrategy() (int, error) {
	// log.Println("randomReplacementStrategy")
	no := rand.Intn(buffersize)
	p := &s.buffer[no]
	if isFixed(p) {
		// log.Printf("%d is fixed", no)

		for n := (no + 1) % buffersize; n != no; n = (n + 1) % buffersize {
			// log.Printf("trying %d", n)
			p = &s.buffer[n]
			if !isFixed(p) {
				return n, nil
			}
		}
		return -1, errors.New("no page to replace")
	}

	return no, nil
}

func lruReplacementStrategy() (int, error) {
	// log.Println("lruReplacementStrategy")
	for {
		if len(s.references) == 0 {
			return -1, errors.New("references exhausted")
		}
		no := s.references[0]
		s.references = s.references[1:]
		p := &s.buffer[no]
		if !isFixed(p) {
			// log.Printf("selected page = %d", no)
			return no, nil
		}
		// log.Printf("lruReplacementStrategy %d is fixed", no)
	}
}

func clockReplacementStrategy() (int, error) {
	// log.Println("clockReplacementStrategy")
	start := s.clock
	for {
		round := 1
		// log.Printf("trying %d counter = %d", s.clock, s.counters[s.clock])
		if s.counters[s.clock] == 0 {
			p := &s.buffer[s.clock]
			if !isFixed(p) {
				// log.Printf("selected page = %d", s.clock)
				s.counters[s.clock] = 0
				return s.clock, nil
			}
		} else {
			s.counters[s.clock] = 0
		}

		s.clock = (s.clock + 1) % buffersize

		if s.clock == start && round == 2 { // no more free pages
			return -1, errors.New("references exhausted")
		}
		round += 1
	}
}
