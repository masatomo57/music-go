package score

import (
	"log"
	"os"

	"github.com/masatomo57/music-go/accompaniment"
	"github.com/masatomo57/music-go/melody"
	"github.com/masatomo57/music-go/util"
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
	"haru_no_umi": {
		Melody:        HaruNoUmiMelody,
		Accompaniment: HaruNoUmiAccompaniment,
	},
}

func (s *Score) WriteToFile(file *os.File) error {
	melody := s.Melody.GenerateSamples()
	accompaniment := s.Accompaniment.GenerateSamples()
	if len(melody) != len(accompaniment) {
		log.Fatalf("length mismatch: melody %d, accompaniment %d", len(melody), len(accompaniment))
	}

	mixed := make([]float32, len(melody))
	for i := range melody {
		mixed[i] = 0.5*melody[i] + 0.5*accompaniment[i]
	}
	
	return util.WriteSamples(file, mixed)
}
