package ingress

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (h *HttpIngressController) handleHttpRequest(w http.ResponseWriter, r *http.Request) {
	os.Stdout.WriteString(fmt.Sprintf("\n\rInbound Http request to %s\n\r", r.RequestURI))
	//TODO: Implement a redirect to https based on the https port configuration

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")

}

func (h *HttpIngressController) handleHttpsRequest(w http.ResponseWriter, r *http.Request) {
	os.Stdout.WriteString(fmt.Sprintf("\n\rInbound Https request to %s\n\r", r.RequestURI))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}
