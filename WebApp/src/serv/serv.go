package serv

import (
    "fmt"
    "net"
    "log"
    "encoding/json"
)

type Message struct {
    Name string	`json:"Name"`
    Request string	`json:"Request"`
    Url string	`json:"Url"`
}

type Reply struct {
	Status string `json:"Status"`
}

func sendReply() {
	conn, err := net.Dial("tcp", "localhost:8020")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := json.NewEncoder(conn)
    r := &Reply{Status: "Success",}
    encoder.Encode(r)
    conn.Close()
    fmt.Printf("sent reply: %s\n-----------\n",r);
}

func handleConnection(conn net.Conn) {
    dec := json.NewDecoder(conn)
    m := Message{}
    dec.Decode(&m)
    fmt.Printf("-----------\nReceived : %s\n", m);
    sendReply()
    conn.Close()
}

func Bootserv() {
    fmt.Println("start compute server");
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}