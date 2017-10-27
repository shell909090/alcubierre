package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var (
	ErrSectionTooBig = errors.New("section too big")
)

const (
	FLAG_PAIR = 0x01
	FLAG_DATA = 0x02
)

type FrameHeader struct {
	Flag    uint8
	Size    uint16
	Section uint16
}

func ReadFrame(r io.Reader) (data []byte, err error) {
	header := &FrameHeader{}
	err = binary.Read(r, binary.BigEndian, header)
	if err != nil {
		return
	}

	data = make([]byte, header.Section)
	_, err = io.ReadFull(r, data)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	return
}

func WriteFrame(w io.Writer, data []byte) (err error) {
	secsize := len(data)
	if secsize > 2^16-1 {
		return ErrSectionTooBig
	}
	var buf bytes.Buffer
	buf.Grow(int(secsize + 5))

	header := &FrameHeader{
		Flag:    FLAG_DATA,
		Size:    uint16(len(data)),
		Section: uint16(secsize),
	}
	binary.Write(&buf, binary.BigEndian, header)
	buf.Write(data)

	b := buf.Bytes()
	n, err := w.Write(b)
	if err != nil {
		return
	}
	if n != len(b) {
		return io.ErrShortWrite
	}
	return
}
