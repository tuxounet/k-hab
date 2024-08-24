package context

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/logger"

	"github.com/urfave/cli/v3"
)

type HabContext struct {
	startContext context.Context
	cwd          string
	config       *config.Config
	log          *logger.Logger
	controllers  map[bases.HabControllers]bases.IController
}

func NewHabContext(startContext context.Context) *HabContext {

	return &HabContext{
		startContext: startContext,
		config:       config.NewConfig(),
		log:          logger.NewLogger(startContext, "Hab"),
	}
}

func (h *HabContext) ParseCli(cmd *cli.Command) error {

	logLevel := cmd.String("loglevel")
	println(logLevel)

	switch logLevel {
	case "TRACE":
		h.log.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		h.log.SetLevel(logrus.DebugLevel)
	case "INFO":
		h.log.SetLevel(logrus.InfoLevel)
	case "WARN":
		h.log.SetLevel(logrus.WarnLevel)
	case "ERROR":
		h.log.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		h.log.SetLevel(logrus.FatalLevel)
	}

	return nil

}

func (h *HabContext) GetHabConfig() bases.HabConfig {
	return h.config.HabConfig
}

func (h *HabContext) SetHabConfig(habConfig bases.HabConfig) {
	h.config.HabConfig = habConfig
}

func (h *HabContext) GetImagesConfig() []bases.HabImageConfig {
	return h.config.ImagesConfig
}

func (h *HabContext) GetContainersConfig() []bases.HabContainerConfig {
	return h.config.ContainersConfig
}

func (h *HabContext) GetLogger() bases.ILogger {
	return h.log
}
func (h *HabContext) GetSubLogger(name string, parent bases.ILogger) bases.ILogger {

	return h.log.CreateSubLogger(name, parent)
}
func (h *HabContext) GetController(controller bases.HabControllers) (bases.IController, error) {
	for key := range h.controllers {
		if key == bases.HabControllers(controller) {
			return h.controllers[key], nil
		}
	}
	return nil, errors.New("controller not found")

}

func (h *HabContext) Getwd() string {
	return h.cwd
}
