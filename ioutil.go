package go_bitmap

import (
	"encoding/binary"
	"io"
)

func ReadByte(r io.Reader) byte {
	var b byte
	ThrowError(binary.Read(r, binary.LittleEndian, &b))
	return b
}

func WriteByte(w io.Writer, b byte) {
	ThrowError(binary.Write(w, binary.LittleEndian, b))
}

func ReadBytes(r io.Reader, length int) []byte {
	b := make([]byte, length)
	_, err := io.ReadFull(r, b)
	ThrowError(err)
	return b
}

func WriteBytes(w io.Writer, data []byte) {
	ThrowError(binary.Write(w, binary.LittleEndian, data))
}

func ReadShortLE(r io.Reader) uint16 {
	var b uint16
	ThrowError(binary.Read(r, binary.LittleEndian, &b))
	return b
}

func WriteShortLE(w io.Writer, b uint16) {
	ThrowError(binary.Write(w, binary.LittleEndian, b))
}

func ReadIntLE(r io.Reader) uint32 {
	var b uint32
	ThrowError(binary.Read(r, binary.LittleEndian, &b))
	return b
}
