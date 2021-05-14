package lzw

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io/ioutil"
)

func Write(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := lzw.NewWriter(buf, lzw.LSB, 8)
	defer writer.Close()
	size, err := writer.Write(data)
	if size != len(data) {
		return nil, fmt.Errorf("Not enough write:%d dataSize:%d", size, len(data))
	}
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Read(data []byte) ([]byte, error) {
	byteReader := bytes.NewReader(data)
	reader := lzw.NewReader(byteReader, lzw.LSB, 8)
	defer reader.Close()

	orgData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return orgData, nil
}
