package pki

import (
	"os"
)

func (p *PKIController) Provision() error {
	p.log.TraceF("Provisioning")

	caPresent, err := p.CAPresent()
	if err != nil {
		return err
	}

	if !caPresent {
		_, err := p.createCA()
		if err != nil {
			return err
		}
	}

	ingressPresent, err := p.IngressCertsPresent()
	if err != nil {
		return err
	}

	if !ingressPresent {
		err := p.createIngressCerts()
		if err != nil {
			return err
		}
	}

	p.log.DebugF("Provisioned")
	return nil
}

func (p *PKIController) Unprovision() error {
	p.log.TraceF("Unprovisioning")

	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return err
	}

	err = os.RemoveAll(pkiPath)
	if err != nil {
		return err
	}

	p.log.DebugF("Unprovisionned")
	return nil
}
