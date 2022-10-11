package net

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net"
)

const (
	endOfMsg = '\n'
)

func WriteConn(conn net.Conn, data []byte) error {
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(data))+1)
	base64.RawURLEncoding.Encode(buf, data)
	buf[len(buf)-1] = endOfMsg
	_, err := conn.Write(buf)
	if err != nil {
		return fmt.Errorf("failed to write to conn: %w", err)
	}
	return nil
}

func ReadConn(conn net.Conn) ([]byte, error) {
	r := bufio.NewReader(conn)
	buf, err := r.ReadBytes(endOfMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to read from conn %w", err)
	}

	b, err := base64.RawURLEncoding.DecodeString(string(buf[:len(buf)-1]))
	if err != nil {
		return nil, fmt.Errorf("failed to decode incoming bytes: %w", err)
	}
	return b, nil
}
