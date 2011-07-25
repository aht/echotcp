package main

import (
	"fmt"
	"flag"
	"net"
	"log"
	"os"
)

var (
	numGoroutine = flag.Int("c", 1024, "number of concurrent goroutines")
	byteCount    = flag.String("n", "", "max number of byte to send")
)

func sendrecv(addr string, done chan bool) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("fatal: ", err)
	}
	defer conn.Close()
	n, err := conn.Write([]byte("I am tcpcc, Hello!"))
	log.Printf("sent %d byte to %s\n", n, conn.RemoteAddr())
	if err != nil {
		log.Println(err)
		done <- false
		return
	}
	buf := make([]byte, n)
	n, err = conn.Read(buf)
	log.Printf("read %d byte from %s\n", n, conn.RemoteAddr())
	if err != nil && err != os.EOF {
		log.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: tcpcc [options..] [host]:[port]")
		os.Exit(-1)
	}
	done := make(chan bool)
	for i := 0; i < *numGoroutine; i++ {
		go sendrecv(flag.Arg(0), done)
	}
	for i := 0; i < *numGoroutine; i++ {
		<-done
	}
}
