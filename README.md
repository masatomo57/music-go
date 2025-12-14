# music-go
Play music using [Go](https://go.dev/).
Inspired by [Golangでカノンを演奏してみた｜homie株式会社](https://note.homie.co.jp/n/n52e3108962d1).

# How to use
- Create and register a score (see [/score](./score)).
- To check your melody or accompaniment, run the following commands:
  - Melody: `sh run.sh <your title> melody` or `sh melody/cmd/run.sh <your title>`.
  - Accompaniment: `sh run.sh <your title> accompaniment` or `sh accompaniment/cmd/run.sh <your title>`.
- Once you are satisfied with the melody and accompaniment, run `sh run.sh <your title>` and enjoy the music!

# Example - Jingle Bells
```sh
sh run.sh
```
