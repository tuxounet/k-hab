package pki

import (
	"fmt"
	"os"
	"path"
	"strings"

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

	profileName := p.ctx.GetConfigValue("hab.name")
	configOrganization := fmt.Sprintf("%s - %s", profileName, p.ctx.GetConfigValue("hab.pki.certs.organization"))
	configOrganizationalUnit := fmt.Sprintf("%s - %s", profileName, p.ctx.GetConfigValue("hab.pki.certs.organization_unit"))
	configCountry := p.ctx.GetConfigValue("hab.pki.certs.country")
	configLocality := p.ctx.GetConfigValue("hab.pki.certs.locality")
	configProvince := p.ctx.GetConfigValue("hab.pki.certs.province")

	configCACommonName := fmt.Sprintf("%s.%s", profileName, p.ctx.GetConfigValue("hab.pki.ca.common_name"))

	rootCAIdentity := goca.Identity{
		Organization:       configOrganization,
		OrganizationalUnit: configOrganizationalUnit,
		Country:            configCountry,
		Locality:           configLocality,
		Province:           configProvince,
		Intermediate:       false,
	}

	RootCA, err := goca.New(configCACommonName, rootCAIdentity)
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

	profileName := p.ctx.GetConfigValue("hab.name")
	configCACommonName := fmt.Sprintf("%s.%s", profileName, p.ctx.GetConfigValue("hab.pki.ca.common_name"))
	os.Setenv("CAPATH", pkiPath)

	RootCA, err := goca.Load(configCACommonName)
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

	profileName := p.ctx.GetConfigValue("hab.name")
	configCACommonName := fmt.Sprintf("%s.%s", profileName, p.ctx.GetConfigValue("hab.pki.ca.common_name"))

	os.Setenv("CAPATH", pkiPath)

	CAs := goca.List()
	for _, ca := range CAs {
		if ca == configCACommonName {
			return true, nil
		}
	}

	return false, nil
}

func (p *PKIController) IngressCertsPresent() (bool, error) {
	pkiPath, err := p.getPKIStoragePath()
	if err != nil {
		return false, err
	}
	os.Setenv("CAPATH", pkiPath)

	profileName := p.ctx.GetConfigValue("hab.name")
	configIngressCommonName := fmt.Sprintf("%s.%s", profileName, p.ctx.GetConfigValue("hab.pki.ingress.common_name"))

	ca, err := p.loadCA()
	if err != nil {
		return false, err
	}

	certs := ca.ListCertificates()

	for _, cert := range certs {
		if cert == configIngressCommonName {
			return true, nil
		}
	}

	return false, nil
}

func (p *PKIController) createIngressCerts() error {

	ca, err := p.loadCA()
	if err != nil {
		return err
	}

	profileName := p.ctx.GetConfigValue("hab.name")
	configOrganization := fmt.Sprintf("%s - %s", profileName, p.ctx.GetConfigValue("hab.pki.certs.organization"))
	configOrganizationalUnit := fmt.Sprintf("%s - %s", profileName, p.ctx.GetConfigValue("hab.pki.certs.organization_unit"))
	configCountry := p.ctx.GetConfigValue("hab.pki.certs.country")
	configLocality := p.ctx.GetConfigValue("hab.pki.certs.locality")
	configProvince := p.ctx.GetConfigValue("hab.pki.certs.province")

	configIngressCommonName := fmt.Sprintf("%s.%s", profileName, p.ctx.GetConfigValue("hab.pki.ingress.common_name"))

	certsDnsNames := p.ctx.GetConfigValue("hab.pki.ingress.dns_names")
	egressCertIdentity := goca.Identity{
		Organization:       configOrganization,
		OrganizationalUnit: configOrganizationalUnit,
		Country:            configCountry,
		Locality:           configLocality,
		Province:           configProvince,
		Intermediate:       false,
		DNSNames:           strings.Split(certsDnsNames, ","),
	}

	_, err = ca.IssueCertificate(configIngressCommonName, egressCertIdentity)
	if err != nil {
		return err
	}
	return nil

}
