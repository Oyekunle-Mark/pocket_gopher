package main

import (
	"io"
	"net"
	"time"
)

var conn net.Conn
var reader io.ReadCloser

func dial(network, address string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}

	netConnection, err := net.DialTimeout(network, address, 5*time.Second)

	if err != nil {
		return nil, err
	}

	return netConnection, nil
}

func closeConnection() {
	if conn != nil {
		conn.Close()
	}

	if reader != nil {
		reader.Close()
	}
}
