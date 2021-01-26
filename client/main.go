package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"shiweikang/pingpongTestClient/model"
)

var (
	ip    string
	size  int
	times int
)

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "set server ip address")
	flag.IntVar(&size, "size", 4*1024*1024, "set send buffer size")
	flag.IntVar(&times, "times", 4, "set send times")
}

func main() {
	flag.Parse()

	client, err := rpc.Dial("tcp", ip+":14567")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply int
	var arg model.BenchArgs
	arg.Size = size
	arg.Times = times

	client.Call("BenchService.Start", arg, &reply)
	if err != nil {
		log.Fatal(err)
	}
	client.Close()
	sendToServer(ip, size, times)

}
func sendToServer(ip string, size int, times int) {
	conn, err := net.Dial("tcp", ip+":14568")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	randomBytes := make([]byte, size, size)
	rand.Read(randomBytes)
	for i := 0; i < times; i++ {
		// n := 0
		_, err = conn.Write(randomBytes)
		if err != nil {
			fmt.Println(err)
		}

		msg := make([]byte, 10)
		_, err := conn.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg))

	}
}
