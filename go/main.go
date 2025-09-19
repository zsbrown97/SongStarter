package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/zsbrown97/songStarter/song"
)

type SongStart struct {
	Chords       string `json:"chords"`
	Instrument   string `json:"instrument"`
	KeySignature string `json:"keySignature"`
	Numerals     string `json:"numerals"`
}

func main() {
	http.HandleFunc("/api/songbones", songBonesHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func songBonesHandler(w http.ResponseWriter, r *http.Request) {
	var instrumentIndex int = rand.Intn(len(Instruments))
	var majorMinor int = rand.Intn(2) // 0 = major, 1 = minor
	var key string = GetKeySignature(majorMinor)
	
	progression := song.ChordProgression(key, 4, majorMinor)

	var chords string = Stringify(progression[0])
	var numerals string = Stringify(progression[1])

	result := SongStart{
		Chords: chords,
		Instrument: Instruments[instrumentIndex],
		KeySignature: key,
		Numerals: numerals,
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
