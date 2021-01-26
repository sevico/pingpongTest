package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	bench := new(BenchService)
	rpc.RegisterName("BenchService", bench)
	dataListener, err := net.Listen("tcp", ":14568")
	listener, err := net.Listen("tcp", ":14567")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		rpc.ServeConn(conn)

		receive(dataListener, bench.Size, bench.ReminTimes)
	}

}

func receive(listener net.Listener, size, times int) {

	fmt.Println(size, times)

	data := make([]byte, size, size)
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for i := 0; i < times; i++ {
		n := 0
		received := 0
		for received = 0; received < size; received += n {
			var err error
			conn.SetDeadline(time.Now().Add(time.Second))
			n, err = conn.Read(data[:size-received])
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Printf("receive %d size bytes\n", received)
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
