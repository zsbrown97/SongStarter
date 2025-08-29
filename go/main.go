package main

import (
	"fmt"
	"math/rand"
)

var instruments = []string {
	"Piano",
	"Guitar",
	"Bass",
	"Synth",
}

func main() {
	majorMinor := rand.Intn(2)
	instrument := rand.Intn(len(instruments))
	randomSongStart(majorMinor, instrument)
}

func randomSongStart(m int, i int) {
	fmt.Println("Key Signature: " + keySignature(m))
	switch m {
	case 0:
		fmt.Println("Here are some chords: " + majorProgression(4))
	case 1:
		fmt.Println("Here are some chords: " + naturalMinorProgression(4))
	}
	fmt.Println("Start with a " + instruments[i])
}


