package cmd

import (
	"github.com/tuxounet/k-hab/hab"
)

func ShellCmd(hab *hab.Hab) error {

	return hab.Shell()
}
