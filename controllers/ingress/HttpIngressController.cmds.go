package ingress

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/pki"
)

func (l *HttpIngressController) getPKIController() (*pki.PKIController, error) {
	controller, err := l.ctx.GetController(bases.PKIController)
	if err != nil {
		return nil, err
	}
	pkiController := controller.(*pki.PKIController)
	return pkiController, nil
}
