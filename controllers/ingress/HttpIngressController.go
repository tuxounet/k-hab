package ingress

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/tuxounet/k-hab/bases"
)

type HttpIngressController struct {
	bases.BaseController
	ctx         bases.IContext
	log         bases.ILogger
	httpServer  *http.Server
	httpsServer *http.Server
	certFile    string
	keyFile     string
}

func NewHttpIngress(ctx bases.IContext) *HttpIngressController {
	return &HttpIngressController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.IngressController), ctx.GetLogger()),
	}
}

func (h *HttpIngressController) Start() error {
	h.log.TraceF("Starting")

	ingress_host := h.ctx.GetConfigValue("hab.ingress.listen.host")
	ingress_port_http := h.ctx.GetConfigValue("hab.ingress.listen.port.http")
	ingress_port_https := h.ctx.GetConfigValue("hab.ingress.listen.port.https")

	h.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", ingress_host, ingress_port_http),
		Handler: http.HandlerFunc(h.handleHttpRequest),
	}

	h.httpsServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", ingress_host, ingress_port_https),
		Handler: http.HandlerFunc(h.handleHttpsRequest),
	}

	pkiController, err := h.getPKIController()
	if err != nil {
		return err
	}

	ingressCertFile, err := pkiController.GetIngressCertFile()
	if err != nil {
		return err
	}
	ingressKeyFile, err := pkiController.GetIngressKeyFile()
	if err != nil {
		return err
	}

	h.certFile = ingressCertFile
	h.keyFile = ingressKeyFile

	go h.listen()

	return nil
}

func (h *HttpIngressController) listen() {
	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := h.httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()

		err := h.httpsServer.ListenAndServeTLS(h.certFile, h.keyFile)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	h.log.DebugF("Listening on %s:%s and %s:%s", h.httpServer.Addr, h.httpsServer.Addr)
	wg.Wait()

}

func (h *HttpIngressController) Stop() error {

	if h.httpServer != nil {
		h.httpServer.Close()
		h.httpServer = nil
	}
	return nil
}
