package main

import (
	"log"

	"github.com/sh3rp/turl/server"
)

func main() {
	s := server.TurlServer{}
	log.Printf("%v\n", s.StartProxy())
}
