package main

import (
	"flag"
	"log"
	"os"

	score "github.com/masatomo57/music-go/score"
)

func main() {
	title := flag.String("title", "jingle_bell", "")
	flag.Parse()

	score, ok := score.Scores[*title]
	if !ok {
		log.Fatalf("score not found: %s", *title)
	}

	file := "out.bin"
	f, _ := os.Create(file)
	defer f.Close()

	score.WriteToFile(f)
}
