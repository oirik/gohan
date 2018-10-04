package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {

	var port = flag.Int("port", 8080, "http server port number")
	var path = flag.String("path", ".", "static files directory")

	var ssl = flag.Bool("ssl", false, "use SSL")
	var certFile = flag.String("certFile", "cert.pem", "SSL cert file path")
	var keyFile = flag.String("keyFile", "key.pem", "SSL key file path")

	var proxy = flag.String("proxy", "", "reverse proxy destination host. ex) localhost:8080")

	flag.Parse()

	fileServer := http.StripPrefix("/", http.FileServer(http.Dir(*path)))

	if *ssl {
		*port = 443
	}
	addr := fmt.Sprintf(":%d", *port)

	pathHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/") && !strings.HasSuffix(r.URL.Path, "/") {
			fileServer.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	proxyHandler := &httputil.ReverseProxy{Director: func(r *http.Request) {
		r.URL.Scheme = "http"
		r.URL.Host = *proxy
		r.Host = *proxy
	}}

	var err error
	if *ssl {
		if *proxy != "" {
			log.Printf("https server starting as reverse proxy. port=%d proxy=%s certFile=%s keyFile=%s\n", *port, *proxy, *certFile, *keyFile)
			err = http.ListenAndServeTLS(addr, *certFile, *keyFile, proxyHandler)
		} else {
			log.Printf("https server starting. port=%d path=%s certFile=%s keyFile=%s\n", *port, *path, *certFile, *keyFile)
			err = http.ListenAndServeTLS(addr, *certFile, *keyFile, pathHandler)
		}
	} else {
		if *proxy != "" {
			log.Printf("http server starting as reverse proxy. port=%d proxy=%s\n", *port, *proxy)
			err = http.ListenAndServe(addr, proxyHandler)
		} else {
			log.Printf("http server starting. port=%d path=%s\n", *port, *path)
			err = http.ListenAndServe(addr, pathHandler)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
