#!/bin/bash
set -e
title=${1:-jingle_bell}

go run main.go -title=$title

output_file="${title}_melody.wav"
ffplay -autoexit "$output_file"
