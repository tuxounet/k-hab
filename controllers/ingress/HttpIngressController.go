package ingress

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tuxounet/k-hab/bases"

	"github.com/tuxounet/k-hab/utils"
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
		log: ctx.GetSubLogger("HttpIngressController", ctx.GetLogger()),
	}
}

func (h *HttpIngressController) handleProxy(w http.ResponseWriter, r *http.Request) {
	os.Stdout.WriteString(fmt.Sprintf("\n\rInbound Http request to %s\n\r", r.RequestURI))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

func (h *HttpIngressController) Start() error {
	h.log.TraceF("Starting")

	config := h.ctx.GetHabConfig()

	ingress_host := utils.GetMapValue(config, "ingress.listen.host").(string)
	ingress_port_http := utils.GetMapValue(config, "ingress.listen.port.http").(string)

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
