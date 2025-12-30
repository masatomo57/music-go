#!/bin/bash
set -e
title=${1:-jingle_bell}
mode=${2:-mix}

go run main.go -title=$title -mode=$mode

output_file="${title}_${mode}.wav"
ffplay -autoexit "$output_file"
