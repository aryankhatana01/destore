package main

import (
	"fmt"

	"github.com/aryankhatana01/destore/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":8080")

	if err := tr.ListenAndAccept(); err != nil {
		fmt.Println("Error listening and accepting: ", err)
	}
	select {} // block forever so that the goroutine can run
}
