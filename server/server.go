package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/emersion/go-imap/backend/memory"
	"github.com/emersion/go-imap/server"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: server cert.pem cert.key")
		return
	}

	// Create a memory backend (you can login with "username:password")
	be := memory.New()

	// https://gist.github.com/spikebike/2232102#file-server-go
	cert, err := tls.LoadX509KeyPair(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}}

	// Create a new server
	s := server.New(be)
	s.TLSConfig = &tlsConfig
	s.Addr = ":993"

	log.Println("Starting IMAP TLS server at :993")
	if err := s.ListenAndServeTLS(); err != nil {
		log.Fatal(err)
	}
}
