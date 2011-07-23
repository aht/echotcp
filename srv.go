package main

import ("net"; "io"; "os"; "log")

func die(err os.Error) {
    if err != nil { log.Fatal(err) }
}

func main() {
    server, err := net.Listen("tcp", "127.0.0.1:3640"); die(err)
    for {
        conn, err := server.Accept(); die(err)
        go func() {
            defer conn.Close()
            n, err := io.Copy(conn, conn)
            log.Printf("echoed %d byte to %s\n", n, conn.RemoteAddr())
            if err != nil {
                log.Printf("error: %s", err)
            }
        }()
    }
}
