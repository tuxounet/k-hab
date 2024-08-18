package cmd

import "github.com/tuxounet/k-hab/hab"

func UnprovisionCmd(hab *hab.Hab) error {

	return hab.Unprovision()
}
