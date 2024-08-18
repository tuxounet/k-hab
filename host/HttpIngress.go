package host

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tuxounet/k-hab/utils"
)

type HttpIngress struct {
	scopeBase string
	habConfig map[string]interface{}
	server    *http.Server
}

func NewHttpIngress(habConfig map[string]interface{}) *HttpIngress {
	return &HttpIngress{
		scopeBase: "HttpIngress",
		habConfig: habConfig,
	}
}

func (h *HttpIngress) handleProxy(w http.ResponseWriter, r *http.Request) {
	os.Stdout.WriteString(fmt.Sprintf("\n\rInbound Http request to %s\n\r", r.RequestURI))

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
}

func (h *HttpIngress) Start(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "Start", func(ctx *utils.ScopeContext) {
		ingress_host := utils.GetMapValue(ctx, h.habConfig, "ingress.listen.host").(string)
		ingress_port_http := utils.GetMapValue(ctx, h.habConfig, "ingress.listen.port.http").(string)

		h.server = &http.Server{
			Addr:    fmt.Sprintf("%s:%s", ingress_host, ingress_port_http),
			Handler: http.HandlerFunc(h.handleProxy),
		}
		go func() {
			err := h.server.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				ctx.Must(err)
			}
		}()

	})
}

func (h *HttpIngress) Stop(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "Stop", func(ctx *utils.ScopeContext) {
		if h.server != nil {
			h.server.Close()
			h.server = nil
		}
	})
}
