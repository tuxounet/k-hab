package cmd

import (
	"github.com/tuxounet/k-hab/hab"
)

func DownCmd(hab *hab.Hab) error {

	return hab.Stop()
}
