package score

import (
	"encoding/binary"
	"log"
	"math"
	"os"

	"github.com/masatomo57/music-go/accompaniment"
	"github.com/masatomo57/music-go/melody"
)

type Score struct {
	Melody        melody.Melody
	Accompaniment accompaniment.Accompaniment
}

var Scores = map[string]Score{
	"jingle_bell": {
		Melody:        JingleBellMelody,
		Accompaniment: JingleBellAccompaniment,
	},
}

func (s Score) WriteToFile(file *os.File) {
	melody := s.Melody.GenerateSamples()
	accompaniment := s.Accompaniment.GenerateSamples()
	if len(melody) != len(accompaniment) {
		log.Fatalf("length mismatch: melody %d, accompaniment %d", len(melody), len(accompaniment))
	}
	samplesLength := len(melody)

	for i := 0; i < samplesLength; i++ {
		sample := 0.5*melody[i] + 0.5*accompaniment[i]
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, math.Float32bits(float32(sample)))
		file.Write(buf)
	}
}
