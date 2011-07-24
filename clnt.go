package main

import (
	"fmt"
	"flag"
	"io"
	"io/ioutil"
	"net"
	"log"
	"os"
)

var (
	numGoroutine = flag.Int("c", 1024, "number of concurrent goroutines")
	filename     = flag.String("f", "", "file to read data from")
	byteCount    = flag.String("n", "", "max number of byte to send")
)

func die(err os.Error) {
	if err != nil {
		log.Fatal(err)
	}
}

func sendrecv(addr string, done chan bool) {
	conn, err := net.Dial("tcp", addr)
	die(err)
	defer conn.Close()
	n, err := conn.Write([]byte("I am tcpcc, Hello!"))
	log.Printf("sent %d byte to %s\n", n, conn.RemoteAddr())
	if err != nil {
		log.Printf("error: %s", err)
	}
	n64, err := io.Copy(ioutil.Discard, conn)
	log.Printf("read %d byte from %s\n", n64, conn.RemoteAddr())
	if err != nil {
		log.Printf("error: %s", err)
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
	for i := 0; i <= *numGoroutine; i++ {
		go sendrecv(flag.Arg(0), done)
	}
	for i := 0; i <= *numGoroutine; i++ {
		<-done
	}
}
