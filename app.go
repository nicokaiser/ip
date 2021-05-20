package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		ip := r.Header.Get("X-Real-Ip")
		if ip == "" {
			rip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err == nil {
				ip = rip
			}
		}
		_, _ = fmt.Fprintln(w, ip);
	})
	http.ListenAndServe(":80", nil)
}
