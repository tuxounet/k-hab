package cmd

import (
	"github.com/tuxounet/k-hab/hab"
)

func NukeCmd(hab *hab.Hab) error {
	return hab.Nuke()

}
