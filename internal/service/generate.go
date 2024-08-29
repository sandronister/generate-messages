package service

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Payload struct {
	Header  uint16
	Command uint8
	Data    []byte
}

func NewPayload(header uint16, command uint8, data []byte) *Payload {
	return &Payload{
		Header:  header,
		Command: command,
		Data:    data,
	}
}

func (p *Payload) ToBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, p.Header)
	if err != nil {
		return nil, fmt.Errorf("failed to write header: %w", err)
	}
	err = binary.Write(buf, binary.BigEndian, p.Command)
	if err != nil {
		return nil, fmt.Errorf("failed to write command: %w", err)
	}
	err = binary.Write(buf, binary.BigEndian, p.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to write data: %w", err)
	}
	return buf.Bytes(), nil
}
