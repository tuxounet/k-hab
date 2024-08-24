package egress

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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
