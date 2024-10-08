package main

import (
	"context"
	_ "embed"
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v3"

	habContext "github.com/tuxounet/k-hab/context"
)

var version = "DEVELOPEMENT"

type Author struct {
	Name    string
	Contact string
}

func main() {

	cmd := &cli.Command{
		Name:                  "k-hab",
		Version:               version,
		Authors:               []any{"Christophe TIRAOUI <github:tuxounet>"},
		Description:           "a single executable configuring and executing one or more containers, whose network interactions are controlled",
		EnableShellCompletion: true,
		HideHelp:              false,
		HideVersion:           false,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "loglevel",
				Value: "INFO",
				Usage: "set minimal level to produce log (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)",
				Validator: func(s string) error {
					switch s {
					case "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL":
						return nil
					default:
						return errors.New("loglevel must be one of TRACE, DEBUG, INFO, WARN, ERROR, FATAL")
					}
				},
			},
			&cli.StringFlag{
				Name:  "setup",
				Value: "",
				Usage: "use a specific setup yaml file, else the built-in one is used",
				Validator: func(s string) error {
					if s == "" {
						return nil
					}
					info, err := os.Stat(s)
					if os.IsNotExist(err) {
						return errors.New("setup file does not exist at " + s)
					}

					if info.IsDir() {
						return errors.New("setup file is a directory and must be a file at " + s)
					}

					return nil
				},
			},
		},
		Commands: []*cli.Command{
			buildCommand("provision", "provision the hab", habContext.ProvisionVerb),
			buildCommand("up", "create and/or launch the hab", habContext.UpVerb),
			buildCommand("start", "create and/or launch the hab", habContext.UpVerb),
			buildCommand("shell", "create and/or launch the hab", habContext.ShellVerb),
			buildCommand("stop", "stop the hab", habContext.DownVerb),
			buildCommand("down", "stop the hab", habContext.DownVerb),
			buildCommand("rm", "rm the hab", habContext.RmVerb),
			buildCommand("unprovision", "unprovision the hab", habContext.UnprovisionVerb),
			buildCommand("nuke", "destroy the hab", habContext.NukeVerb),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}

func buildCommand(name string, usage string, verb habContext.HabVerbs) *cli.Command {
	return &cli.Command{
		Name:  name,
		Usage: usage,
		Action: func(ctx context.Context, ocmd *cli.Command) error {
			workingFolder, err := os.Getwd()
			if err != nil {
				return err
			}

			habCtx := habContext.NewHabContext(ctx, workingFolder)

			logLevel := ocmd.String("loglevel")

			err = habCtx.SetLogLevel(logLevel)
			if err != nil {
				return err
			}

			setup := ocmd.String("setup")

			err = habCtx.SetSetup(setup)
			if err != nil {
				return err
			}

			err = habCtx.Init()
			if err != nil {
				return err
			}

			switch verb {
			case habContext.ProvisionVerb:
				return habCtx.Provision()
			case habContext.UpVerb:
				return habCtx.Start()
			case habContext.ShellVerb:
				return habCtx.Shell()
			case habContext.DownVerb:
				return habCtx.Stop()
			case habContext.RmVerb:
				return habCtx.Rm()
			case habContext.UnprovisionVerb:
				return habCtx.Unprovision()
			case habContext.NukeVerb:
				return habCtx.Nuke()
			default:
				return errors.New("unknown verb " + string(verb))
			}
		},
	}
}
