package util

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/masatomo57/music-go/conf"
)

// wavHeader represents the complete WAV file header structure
type wavHeader struct {
	// RIFFチャンク
	RiffID   [4]byte // "RIFF"
	FileSize uint32  // ファイル全体のサイズ - 8
	WaveID   [4]byte // "WAVE"

	// fmtチャンク
	FmtID         [4]byte // "fmt "
	FmtSize       uint32  // fmtチャンクのサイズ (16)
	Format        uint16  // オーディオフォーマット (3 = IEEE float)
	Channels      uint16  // チャンネル数 (1 = モノラル)
	SampleRate    uint32  // サンプリングレート
	ByteRate      uint32  // 1秒あたりのバイト数
	BlockAlign    uint16  // ブロックアライメント
	BitsPerSample uint16  // 1サンプルあたりのビット数

	// dataチャンク
	DataID   [4]byte // "data"
	DataSize uint32  // オーディオデータのサイズ
}

// WriteWAVHeader writes a WAV header for the given sample count
func WriteWAVHeader(w io.Writer, sampleCount int) error {
	const (
		audioFormat   = 3  // IEEE float
		numChannels   = 1  // モノラル
		bitsPerSample = 32 // float32
	)

	sampleRate := uint32(conf.SamplesPerSec)
	byteRate := sampleRate * numChannels * bitsPerSample / 8
	blockAlign := numChannels * bitsPerSample / 8
	dataSize := uint32(sampleCount * 4) // float32は4バイト

	// ヘッダー構造体を作成
	header := wavHeader{
		RiffID:   [4]byte{'R', 'I', 'F', 'F'},
		FileSize: 36 + dataSize, // 36はヘッダーの固定サイズ
		WaveID:   [4]byte{'W', 'A', 'V', 'E'},

		FmtID:         [4]byte{'f', 'm', 't', ' '},
		FmtSize:       16,
		Format:        audioFormat,
		Channels:      numChannels,
		SampleRate:    sampleRate,
		ByteRate:      byteRate,
		BlockAlign:    uint16(blockAlign),
		BitsPerSample: bitsPerSample,

		DataID:   [4]byte{'d', 'a', 't', 'a'},
		DataSize: dataSize,
	}

	// 構造体を一度にバイナリ書き込み
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, header); err != nil {
		return err
	}

	// バッファの内容をwriterに書き込み
	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}