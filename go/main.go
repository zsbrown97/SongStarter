package main

import (
	"fmt"
	"math/rand"

	"github.com/zsbrown97/songStarter/song"
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
	fmt.Println("Key Signature: " + song.KeySignature(m))
	switch m {
	case 0:
		fmt.Println("Here are some chords: " + song.MajorProgression(4))
	case 1:
		fmt.Println("Here are some chords: " + song.NaturalMinorProgression(4))
	}
	fmt.Println("Start with a " + instruments[i])
}


