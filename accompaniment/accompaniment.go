package accompaniment

import (
	conf "github.com/masatomo57/music-go/conf"
	"github.com/masatomo57/music-go/util"
)

type Accompaniment []struct {
	conf.Chord
	Length float64
}

// GenerateSamples generates audio samples as []float32 without writing to file
func (a *Accompaniment) GenerateSamples() []float32 {
	var samples []float32
	for _, chordWithLength := range *a {
		// 和音の周波数リストを作成
		frequencies := make([]float64, len(chordWithLength.Chord.Notes))
		for i, note := range chordWithLength.Chord.Notes {
			frequencies[i] = note.Hertz()
		}

		chordSamples := util.GenerateSamplesForLength(chordWithLength.Length, func(i int) float64 {
			if len(frequencies) == 0 {
				return 0
			}
			// 各音の振幅を均等に分配
			vol := 1.0 / float64(len(frequencies))
			var sample float64
			for _, freq := range frequencies {
				sample += vol * util.GenerateSineWave(freq, i)
			}
			return sample
		})
		samples = append(samples, chordSamples...)
	}
	return samples
}
