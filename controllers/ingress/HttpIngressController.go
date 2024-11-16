package ingress

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tuxounet/k-hab/bases"
)

type HttpIngressController struct {
	bases.BaseController
	ctx    bases.IContext
	log    bases.ILogger
	server *http.Server
}

func NewHttpIngress(ctx bases.IContext) *HttpIngressController {
	return &HttpIngressController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.IngressController), ctx.GetLogger()),
	}
}

func (h *HttpIngressController) handleProxy(w http.ResponseWriter, r *http.Request) {
	os.Stdout.WriteString(fmt.Sprintf("\n\rInbound Http request to %s\n\r", r.RequestURI))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

func (h *HttpIngressController) Start() error {
	h.log.TraceF("Starting")

	ingress_host := h.ctx.GetConfigValue("hab.ingress.listen.host")
	ingress_port_http := h.ctx.GetConfigValue("hab.ingress.listen.port.http")

	h.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", ingress_host, ingress_port_http),
		Handler: http.HandlerFunc(h.handleProxy),
	}
	go func() {
		err := h.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	h.log.DebugF("Started")
	return nil
}

func (h *HttpIngressController) Stop() error {

	if h.server != nil {
		h.server.Close()
		h.server = nil
	}
	return nil
}
