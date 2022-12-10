package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", pong)
	log.Fatal(http.ListenAndServe(":8840", nil))

}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func pong(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k, v := range header {
		// 将请求中的header挨个写入响应
		w.Header().Set(k, v[0])
	}
	os.Setenv("VERSION", "v1.1.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	log.Printf("ip: %s; code: %d\n", r.RemoteAddr, http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
