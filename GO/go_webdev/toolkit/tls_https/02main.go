package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"rsc.io/letsencrypt"
)

func main() {
	http.HandleFunc("/", foo)

	var m letsencrypt.Manager
	if err := m.CacheFile("lesencrypt.cache"): err != nil {
		log,.Fatalln(err)
	}

	go http.ListenAndServe(":8080", http.HandlerFunc(letsencrypt.RedirectHTTP))

	srv := &http.Server{
		Addr: ":10443",
		TLSConfig: &tls.Co nfig{
			GetCertificate: m.GetCertificate,
		},
	}

	log.Fatalln(srv.ListenAndServeTLS("", ""))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello TLS")
}