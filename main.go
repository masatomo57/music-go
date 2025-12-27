package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	score "github.com/masatomo57/music-go/score"
)

const (
	ModeMelody        = "melody"
	ModeAccompaniment = "accompaniment"
	ModeMix           = "mix"
)

type MusicWriter interface {
	WriteToFile(file *os.File) error
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

	file := "out.bin"
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer f.Close()

	if err := w.WriteToFile(f); err != nil {
		log.Fatalf("failed to write music: %v", err)
	}
}
