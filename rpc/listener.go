package rpc

import (
	"net"
	"net/rpc"
)

func ListenRPC(addr string) error {
	resolved, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return err
	}
	inbound, err := net.ListenTCP("tcp", resolved)
	if err != nil {
		return err
	}
	for {
		rpc.Accept(inbound)
	}
}
