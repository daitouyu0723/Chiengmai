package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var data []byte

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", ":3306")
	checkError(err)
	defer conn.Close()
	data = make([]byte, 4)
	br := bufio.NewReader(conn)
	data, err = br.Peek(4)
	checkError(err)
	br.Discard(4)
	size := int(data[0]) + int(data[1]<<8) + int(data[2]<<16)
	//sequenctId := int(data[4])

	payload, err := br.Peek(size)
	checkError(err)
	fmt.Println(string(payload))
}
