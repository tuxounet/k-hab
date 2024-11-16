package egress

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/tuxounet/k-hab/bases"
)

type HttpEgressController struct {
	bases.BaseController
	ctx    bases.IContext
	log    bases.ILogger
	server *http.Server
}

func NewHttpEgressController(ctx bases.IContext) *HttpEgressController {

	return &HttpEgressController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.EgressController), ctx.GetLogger()),
	}
}

func (h *HttpEgressController) Start() error {

	egress_host := h.ctx.GetConfigValue("hab.incus.host.address")
	egress_port := h.ctx.GetConfigValue("hab.egress.listen.port")

	h.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", egress_host, egress_port),
		Handler: http.HandlerFunc(h.handleProxy),
		// Disable HTTP/2.
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	go func() {
		err := h.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return nil
}

func (h *HttpEgressController) Stop() error {

	if h.server != nil {
		h.server.Close()
		h.server = nil
	}

	return nil
}
