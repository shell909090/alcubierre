package main

import "net"

type MultiLink struct {
	queue chan []byte
	conns []net.Conn
	udp   *net.UDPConn
}

func NewMultiLink(udp *net.UDPConn) (m *MultiLink) {
	m = &MultiLink{
		queue: make(chan []byte, 16),
		udp:   udp,
	}
	return m
}

func (m *MultiLink) ReadFromUdp() {
	for {
		b := make([]byte, 8192)
		n, err := m.udp.Read(b)
		if err != nil {
			// TODO:
			return
		}

		select {
		case c <- b[:n]:
		default:
		}
	}
}

func (m *MultiLink) Add(c net.Conn) {
	m.conns = append(m.conns, c)
	go m.ReadFromConn(c)
	go m.WriteToConn(c)
}

func (m *MultiLink) ReadFromConn(c net.Conn) {
	for {
		data, err = ReadFrame(c)
		if err != nil {
			// TODO:
			return
		}
		m.udp.Write(buf)
	}
}

func (m *MultiLink) WriteToConn(c net.Conn) {
	for {
		b := <-c
		err = WriteFrame(c, b)
		if err != nil {
			// TODO:
			return
		}
	}
}

func RunServer(cfg *Config) (err error) {
	return
}
