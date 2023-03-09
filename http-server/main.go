package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"httpserver/metrics"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	// 指标汇报
	metrics.Register()
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/pong", pong)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8840", nil))

}

func healthz(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal() // 这个函数执行最后记录一次指标
	// 随机sleep几秒
	randomSeconds := rand.Intn(5) + 1
	duration := time.Duration(randomSeconds) * time.Second
	time.Sleep(duration)

	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func pong(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal() // 这个函数执行最后记录一次指标
	// 随机sleep几秒
	randomSeconds := randInt(1, 20)
	duration := time.Duration(randomSeconds) * time.Second
	time.Sleep(duration)

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

// 随机整数秒
func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
