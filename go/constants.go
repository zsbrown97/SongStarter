package main

import (
	"math/rand"
	"strings"

	"github.com/zsbrown97/songStarter/song"
)

// Functions
func GetKeySignature(majorMinor int) string {
	keyIndex := rand.Intn(len(song.KeySignatures))
	return song.KeySignatures[keyIndex][majorMinor]
}

func Stringify(progression []string) string {
	return strings.Join(progression, " ")
}

// Slices
var Instruments = []string {
	"Bass",
	"Guitar",
	"Piano",
	"Synth",
}