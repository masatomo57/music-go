package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	score "github.com/masatomo57/music-go/score"
	"github.com/masatomo57/music-go/util"
)

func main() {
	title := flag.String("title", "jingle_bell", "")
	flag.Parse()

	score, ok := score.Scores[*title]
	if !ok {
		log.Fatalf("score not found: %s", *title)
	}
	a := score.Accompaniment

	file := fmt.Sprintf("%s_accompaniment.wav", *title)
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	defer f.Close()

	// サンプルを生成
	samples := a.GenerateSamples()

	// WAVヘッダーを書き込み
	if err := util.WriteWAVHeader(f, len(samples)); err != nil {
		log.Fatalf("failed to write WAV header: %v", err)
	}

	// 音声データを書き込み
	if err := util.WriteSamples(f, samples); err != nil {
		log.Fatalf("failed to write samples: %v", err)
	}

	fmt.Printf("Generated %s for %s\n", file, *title)
}
