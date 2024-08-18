package cmd

import (
	"github.com/tuxounet/k-hab/hab"
)

func UpCmd(hab *hab.Hab) error {

	return hab.Start()
}
