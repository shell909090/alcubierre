package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	data := []byte("\x01\x02")

	err := WriteFrame(&buf, data)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	fmt.Println("%v", buf.Bytes())

	data, err = ReadFrame(&buf)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	fmt.Println("%v", data)

	return
}
