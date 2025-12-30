package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	score "github.com/masatomo57/music-go/score"
	"github.com/masatomo57/music-go/util"
)

const (
	ModeMelody        = "melody"
	ModeAccompaniment = "accompaniment"
	ModeMix           = "mix"
)

type MusicWriter interface {
	GenerateSamples() []float32
}

func getMusicWriter(mode string, score *score.Score) (MusicWriter, error) {
	switch mode {
	case ModeMelody:
		return &score.Melody, nil
	case ModeAccompaniment:
		return &score.Accompaniment, nil
	case ModeMix:
		return score, nil
	default:
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}

}

func main() {
	title := flag.String("title", "jingle_bell", "")
	mode := flag.String("mode", "mix", "one of (melody | accompaniment | mix)")

	flag.Parse()

	score, ok := score.Scores[*title]
	if !ok {
		log.Fatalf("score not found: %s", *title)
	}

	w, err := getMusicWriter(*mode, &score)
	if err != nil {
		log.Fatalf("failed to get music writer: %v", err)
	}

	file := fmt.Sprintf("%s_%s.wav", *title, *mode)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer f.Close()

	// サンプルを生成
	samples := w.GenerateSamples()

	// WAVヘッダーを書き込み
	if err := util.WriteWAVHeader(f, len(samples)); err != nil {
		log.Fatalf("failed to write WAV header: %v", err)
	}

	// 音声データを書き込み
	if err := util.WriteSamples(f, samples); err != nil {
		log.Fatalf("failed to write samples: %v", err)
	}

	fmt.Printf("Generated %s (%s mode)\n", file, *mode)
}
