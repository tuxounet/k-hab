package egress

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

func (h *HttpEgressController) handleProxy(w http.ResponseWriter, r *http.Request) {

	os.Stdout.WriteString(fmt.Sprintf("\nIncoming request to %s\n\r", r.RequestURI))

	if r.Method == http.MethodConnect {
		h.handleTunneling(w, r)
	} else {
		h.handleHttpProxy(w, r)
	}

}

func (h *HttpEgressController) handleHttpProxy(w http.ResponseWriter, r *http.Request) {
	os.Stdout.WriteString(fmt.Sprintf("\n\rProxying request to %s\n\r", r.RequestURI))

	resp, err := http.Get(r.RequestURI)

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("\nError with proxy request to %s: %s\n\r", r.RequestURI, err.Error()))
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for key, value := range resp.Header {
		w.Header().Set(key, value[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func (h *HttpEgressController) handleTunneling(w http.ResponseWriter, r *http.Request) {
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}
	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}
