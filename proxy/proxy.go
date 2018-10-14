package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"./go-imap-proxy"

	"github.com/emersion/go-imap/server"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage: proxy imap.example.com:993 cert.pem cert.key")
		return
	}

	// https://gist.github.com/spikebike/2232102#file-server-go
	cert, err := tls.LoadX509KeyPair(os.Args[2], os.Args[3])
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}}

	be := proxy.NewTLS(os.Args[1], nil)

	// Create a new server
	s := server.New(be)
	s.TLSConfig = &tlsConfig
	s.Addr = ":993"

	log.Printf("Starting IMAP TLS server at :993 proxying to %s\n", os.Args[1])
	if err := s.ListenAndServeTLS(); err != nil {
		log.Fatal(err)
	}
}
