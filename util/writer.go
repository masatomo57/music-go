package util

import (
	"encoding/binary"
	"io"
	"math"
)

// WriteSamples writes float32 samples to a writer in little-endian format
func WriteSamples(w io.Writer, samples []float32) error {
	buf := make([]byte, 4)
	for _, sample := range samples {
		binary.LittleEndian.PutUint32(buf, math.Float32bits(sample))
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}
