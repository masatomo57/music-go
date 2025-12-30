package util

import (
	"math"

	"github.com/masatomo57/music-go/conf"
)

// FrequencyProvider provides frequency value for sound generation
type FrequencyProvider func(i int) float64

// GenerateSamplesForLength generates audio samples for a given length with custom frequency calculation
func GenerateSamplesForLength(length float64, freqProvider FrequencyProvider) []float32 {
	sampleCount := int((length * conf.SamplesPerSec) / 4)
	damping := math.Pow(conf.End, 1.0/float64(sampleCount))

	samples := make([]float32, sampleCount)
	for i := 0; i < sampleCount; i++ {
		sample := freqProvider(i)
		sample = sample * math.Pow(damping, float64(i))
		samples[i] = float32(sample)
	}
	return samples
}

// GenerateSineWave generates a sine wave at the given frequency
func GenerateSineWave(frequency float64, i int) float64 {
	return math.Sin((2 * math.Pi * frequency * float64(i)) / float64(conf.SamplesPerSec))
}
