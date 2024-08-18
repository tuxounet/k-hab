package cmd

import (
	"github.com/tuxounet/k-hab/hab"
)

func ProvisionCmd(hab *hab.Hab) error {

	return hab.Provision()
}
