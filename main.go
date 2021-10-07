package main

import (
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	header := req.Header
	for k, values := range header {
		for _, v := range values {
			w.Header().Set(k, v)
		}
	}

	vsersion := os.Getenv("VERSION")
	w.Header().Set("VERSION", vsersion)
	ip := RemoteIP(req)
	w.WriteHeader(http.StatusOK)
	log.Printf("client ip: %v, statusCode: %v", ip, http.StatusOK)
}

func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
