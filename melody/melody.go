package melody

import (
	conf "github.com/masatomo57/music-go/conf"
	"github.com/masatomo57/music-go/util"
)

type Melody []struct {
	conf.Note
	Length float64
}

// GenerateSamples generates audio samples as []float32 without writing to file
func (m *Melody) GenerateSamples() []float32 {
	var samples []float32
	for _, noteWithLength := range *m {
		frequency := noteWithLength.Note.Hertz()
		noteSamples := util.GenerateSamplesForLength(noteWithLength.Length, func(i int) float64 {
			return util.GenerateSineWave(frequency, i)
		})
		samples = append(samples, noteSamples...)
	}
	return samples
}
