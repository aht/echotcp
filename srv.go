package main

import ("net"; "io"; "os"; "log")

func die(err os.Error) {
    if err != nil { log.Fatal(err) }
}

func main() {
    server, err := net.Listen("tcp", "0.0.0.0:3640"); die(err)
    for {
        conn, err := server.Accept();
        if err != nil {
            log.Printf("error: %s", err)
            continue
        }
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
