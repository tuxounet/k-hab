package egress

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tuxounet/k-hab/bases"

	"github.com/tuxounet/k-hab/utils"
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

func (h *HttpEgressController) Start() error {

	egress_host := utils.GetMapValue(h.ctx.GetHabConfig(), "lxd.lxc.host.address").(string)
	egress_port := utils.GetMapValue(h.ctx.GetHabConfig(), "egress.listen.port").(string)

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
