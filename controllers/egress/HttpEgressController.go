package egress

import (
	"fmt"
	"net/http"
	"os"

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
		log: ctx.GetSubLogger("HttpEgressController", ctx.GetLogger()),
	}
}

func (h *HttpEgressController) handleProxy(w http.ResponseWriter, r *http.Request) {

	os.Stdout.WriteString(fmt.Sprintf("\nIncoming request to %s\n\r", r.RequestURI))

	h.handleHttpProxy(w, r)

}

func (h *HttpEgressController) Start() error {

	egress_host := h.ctx.GetConfigValue("hab.lxd.lxc.host.address")
	egress_port := h.ctx.GetConfigValue("hab.egress.listen.port")

	h.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", egress_host, egress_port),
		Handler: http.HandlerFunc(h.handleProxy),
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
