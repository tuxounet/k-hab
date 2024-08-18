package host

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/utils"
)

type HttpEgress struct {
	scopeBase string
	habConfig config.HabConfig
	server    *http.Server
	ctx       *utils.ScopeContext
}

func NewHttpEgress(habConfig config.HabConfig) *HttpEgress {

	return &HttpEgress{
		scopeBase: "HttpEgress",
		habConfig: habConfig,
	}
}

func (h *HttpEgress) handleProxy(w http.ResponseWriter, r *http.Request) {

	os.Stdout.WriteString(fmt.Sprintf("\nIncoming request to %s\n\r", r.RequestURI))
	uri, err := url.Parse(r.RequestURI)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("\nError with request to %s: %s\n\r", r.RequestURI, err.Error()))
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
	if uri.IsAbs() {
		h.handleHttpProxy(w, r)
	} else {
		if strings.HasPrefix(uri.Path, "/os/") {
			h.handleOsMirror(w, r)
		} else {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}

}

func (h *HttpEgress) Start(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "Start", func(ctx *utils.ScopeContext) {
		egress_host := utils.GetMapValue(ctx, h.habConfig, "lxd.lxc.host.address").(string)
		egress_port := utils.GetMapValue(ctx, h.habConfig, "egress.listen.port").(string)

		h.ctx = ctx

		h.server = &http.Server{
			Addr:    fmt.Sprintf("%s:%s", egress_host, egress_port),
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

func (h *HttpEgress) Stop(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "Stop", func(ctx *utils.ScopeContext) {
		if h.server != nil {
			h.server.Close()
			h.server = nil
		}
	})
}
