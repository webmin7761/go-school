package main

import (
	"bufio"
	"fmt"
	"net"
)

func InitTCP(addr string) (err error) {
	var (
		listener net.Listener
	)

	if listener, err = net.Listen("tcp", addr); err != nil {
		fmt.Printf("net.Listen(tcp, %s) error(%v)", addr, err)
		return
	}

	acceptTCP(listener)

	return
}

func acceptTCP(lis net.Listener) {
	var (
		conn net.Conn
		err  error
	)
	for {

		if conn, err = lis.Accept(); err != nil {
			fmt.Printf("listener.Accept(\"%s\") error(%v)", lis.Addr().String(), err)
			return
		}

		go serveTCP(conn)
	}
}

func serveTCP(conn net.Conn) {
	defer conn.Close()
	p := &Proto{}
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {

		err := p.ReadTCP(rd)
		if err != nil {
			fmt.Printf("read error [%v]", err)
			return
		}

		fmt.Printf("seq:[%d] op:[%d] body:[%s]\n", p.Seq, p.Op, string(p.Body))

		err = p.WriteTCP(wr)
		if err != nil {
			fmt.Printf("write error [%v]", err)
			return
		}
	}
}
