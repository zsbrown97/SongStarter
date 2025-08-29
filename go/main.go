package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/zsbrown97/songStarter/song"
)

type SongStart struct {
	KeySignature string `json:"keySignature"`
	Chords string `json:"chords"`
	Instrument string `json:"instrument"`
}

var instruments = []string {
	"Piano",
	"Guitar",
	"Bass",
	"Synth",
}

func main() {
	http.HandleFunc("/api/songstart", songStartHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func songStartHandler(w http.ResponseWriter, r *http.Request) {
	majorMinor := rand.Intn(2)
	instrument := rand.Intn(len(instruments))

	var chords string

	switch majorMinor {
	case 0:
		chords = song.MajorProgression(4)
	case 1:
		chords = song.NaturalMinorProgression(4)
	}

	result := SongStart{
		KeySignature: song.KeySignature(majorMinor),
		Chords: chords,
		Instrument: instruments[instrument],
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
