package score

import (
	accompaniment "github.com/masatomo57/music-go/accompaniment"
	conf "github.com/masatomo57/music-go/conf"
	melody "github.com/masatomo57/music-go/melody"
)

// 春の海 (Haru no Umi) - 宮城道雄
// 尺八パートの有名な冒頭旋律
var HaruNoUmiMelody = melody.Melody{
	{Note: conf.E4, Length: 6},
	{Note: conf.B4, Length: 1},
	{Note: conf.D5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.A5, Length: 1},
	{Note: conf.B5, Length: 1},
	{Note: conf.A5, Length: 2},
	{Note: conf.E5, Length: 11},

	{Note: conf.E4, Length: 6},
	{Note: conf.B4, Length: 1},
	{Note: conf.D5, Length: 1},
	{Note: conf.E5, Length: 1},
	{Note: conf.A5, Length: 1},
	{Note: conf.B5, Length: 1},
	{Note: conf.D6, Length: 2},
	{Note: conf.E6, Length: 11},
}

var HaruNoUmiAccompaniment = accompaniment.Accompaniment{
	{Chord: conf.ChordNone, Length: 24},

	{Chord: conf.ChordNone, Length: 24},
}