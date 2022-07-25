// 作业： Oliver  学号：G20220797070039
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

/*
1. 接收客户端 request，并将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 当访问 {url}/healthz 时，应返回200
*/

func index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1>Welcome to oliver testing website<h1>"))

	//2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "v1.0")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	//1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		//fmt.Println(k, v)
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s\n", k, vv)
			w.Header().Set(k, vv)
		}
	}

	clientIP := getClientIP(r)
	//fmt.Println(r.RemoteAddr)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientIP)
	fmt.Fprintf(w, "Success! Client IP:  %s", clientIP)

}

func getClientIP(r *http.Request) string {

	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		fmt.Println("get ip from X-Forwarded-For")
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		fmt.Println("get ip from X-Real-Ip")
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		fmt.Println("get ip from RemoteAddr")
		return ip
	}
	return ""

}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "wokring, response code is 200")
}

func main() {
	mux := http.NewServeMux()
	// 06. debug
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Start http server failed, err: %s\n", err.Error())
	}

}
