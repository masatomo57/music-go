#!/bin/bash
set -e
title=${1:-jingle_bell}
mode=${2:-mix}

go run main.go -title=$title -mode=$mode

ffplay -autoexit -f f32le -showmode 1 out.bin
