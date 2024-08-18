package cmd

import (
	"github.com/tuxounet/k-hab/hab"
)

func RmCmd(hab *hab.Hab) error {

	return hab.Rm()
}
