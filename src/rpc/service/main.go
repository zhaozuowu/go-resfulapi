package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (helloSerivce *HelloService) Hello(request string, replay *string) error {

	*replay = "hello:" + request
	return nil
}


func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
