package pki

import (
	"os"
)

func (p *PKIController) Install() error {
	p.log.TraceF("Installing")

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

	p.log.DebugF("Installed")
	return nil
}

func (p *PKIController) Unintall() error {
	p.log.TraceF("Uninstalling")

	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return err
	}

	err = os.RemoveAll(pkiPath)
	if err != nil {
		return err
	}

	p.log.DebugF("Uninstalled")
	return nil
}
