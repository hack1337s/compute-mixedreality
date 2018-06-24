package client

import (
    "fmt"
    "log"
    "net"
    "encoding/json"
    "sync"
)

var wg sync.WaitGroup

type Message struct {
    Name string	`json:"Name"`
    Request string	`json:"Request"`
    Url string	`json:"Url"`
}

type Reply struct {
	Status string `json:"Status"`
}

func handleMessage(conn net.Conn) {
    dec := json.NewDecoder(conn)
    r := Reply{}
    dec.Decode(&r)
    fmt.Printf("Received : %s\n-----------\n", r);
    conn.Close()
    wg.Done()
}

func listenOverTCP() {
	fmt.Println("listening...");
    ln, err := net.Listen("tcp", "localhost:8020")
    if err != nil {
        // handle error
    }
    for {
        conn, err := ln.Accept() // this blocks until connection or error
        if err != nil {
            // handle error
            continue
        }
        go handleMessage(conn) // a goroutine handles conn so that the loop can accept other connections
    }
}

func sendOverTCP() {
	conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error", err)
    }
    encoder := json.NewEncoder(conn)
    m := Message{Name: "Alex", Request: "Post", Url: "http://"}
    encoder.Encode(m)
    conn.Close()
    fmt.Printf("-----------\nsent message:%s\n", m);
    wg.Add(1)
}

func Bootclient() {
    fmt.Println("start client");
    go listenOverTCP()
    sendOverTCP()
    wg.Wait()
}