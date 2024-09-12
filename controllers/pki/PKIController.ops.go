package pki

import (
	"os"
	"path"

	"github.com/kairoaraujo/goca"
)

func (p *PKIController) getPKIStoragePath() (string, error) {

	root, err := p.ctx.GetStorageRoot()
	if err != nil {
		return "", err
	}
	pkiFolder := p.ctx.GetConfigValue("hab.pki.path")

	pkiPath := path.Join(root, pkiFolder)
	if _, err := os.Stat(pkiPath); os.IsNotExist(err) {
		err := os.MkdirAll(pkiPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return pkiPath, nil

}

func (p *PKIController) createCA() (*goca.CA, error) {
	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return nil, err
	}
	os.Setenv("CAPATH", pkiPath)
	rootCAIdentity := goca.Identity{
		Organization:       p.ctx.GetConfigValue("hab.pki.ca.organization"),
		OrganizationalUnit: p.ctx.GetConfigValue("hab.pki.ca.organization_unit"),
		Country:            p.ctx.GetConfigValue("hab.pki.ca.country"),
		Locality:           p.ctx.GetConfigValue("hab.pki.ca.locality"),
		Province:           p.ctx.GetConfigValue("hab.pki.ca.province"),
		Intermediate:       false,
	}

	RootCA, err := goca.New(p.ctx.GetConfigValue("hab.pki.ca.common_name"), rootCAIdentity)
	if err != nil {
		p.log.ErrorF("Error getting CA: %s", err)
		return nil, err
	}
	return &RootCA, nil
}

func (p *PKIController) loadCA() (*goca.CA, error) {
	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return nil, err
	}
	os.Setenv("CAPATH", pkiPath)

	RootCA, err := goca.Load(p.ctx.GetConfigValue("hab.pki.ca.common_name"))
	if err != nil {
		p.log.ErrorF("Error getting CA: %s", err)
		return nil, err
	}
	return &RootCA, nil
}

func (p *PKIController) CAPresent() (bool, error) {
	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return false, err
	}
	os.Setenv("CAPATH", pkiPath)
	cn := p.ctx.GetConfigValue("hab.pki.ca.common_name")
	CAs := goca.List()
	for _, ca := range CAs {
		if ca == cn {
			return true, nil
		}
	}

	return false, nil
}

func (p *PKIController) EgressCertsPresent() (bool, error) {
	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return false, err
	}
	os.Setenv("CAPATH", pkiPath)
	ca, err := p.loadCA()
	if err != nil {
		return false, err
	}

	certs := ca.ListCertificates()
	certCN := p.ctx.GetConfigValue("hab.pki.certs.egress.common_name")
	for _, cert := range certs {
		if cert == certCN {
			return true, nil
		}
	}

	return false, nil
}

func (p *PKIController) createEgressCerts() error {

	ca, err := p.loadCA()
	if err != nil {
		return err
	}
	certCN := p.ctx.GetConfigValue("hab.pki.certs.egress.common_name")
	egressCertIdentity := goca.Identity{
		Organization:       p.ctx.GetConfigValue("hab.pki.certs.egress.organization"),
		OrganizationalUnit: p.ctx.GetConfigValue("hab.pki.certs.egress.organization_unit"),
		Country:            p.ctx.GetConfigValue("hab.pki.certs.egress.country"),
		Locality:           p.ctx.GetConfigValue("hab.pki.certs.egress.locality"),
		Province:           p.ctx.GetConfigValue("hab.pki.certs.egress.province"),
		Intermediate:       false,
		DNSNames:           []string{"w3.intranet.example.com", "www.intranet.example.com"},
	}
	_, err = ca.IssueCertificate(certCN, egressCertIdentity)
	if err != nil {
		return err
	}
	return nil

}
