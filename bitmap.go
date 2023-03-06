package go_bitmap

import (
	"bytes"
	"github.com/GoFeGroup/go-bitmap/glog"
	"image"
	"image/png"
)

import "C"

type Option struct {
	Width       int
	Height      int
	BitPerPixel int
	Data        []byte
	LogLevel    glog.LEVEL
}

type BitMap struct {
	image image.Image
}

func (m *BitMap) ToPng() []byte {
	buf := new(bytes.Buffer)
	ThrowError(png.Encode(buf, m.image))
	return buf.Bytes()
}

func NewBitMapFromRDP6(option *Option) *BitMap {
	glog.SetLevel(option.LogLevel)
	return (&BitMap{}).LoadRDP60(option)
}

func NewBitmapFromRLE(option *Option) *BitMap {
	glog.SetLevel(option.LogLevel)
	return (&BitMap{}).LoadRLE(option)
}
