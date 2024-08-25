package context

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/logger"
	"github.com/tuxounet/k-hab/context/setup"

	"github.com/urfave/cli/v3"
)

type HabContext struct {
	startContext context.Context
	cwd          string
	config       *config.Config
	setup        *setup.Setup
	log          *logger.Logger
	controllers  map[bases.HabControllers]bases.IController
}

func NewHabContext(startContext context.Context, defaultConfig map[string]string, defaultSetup bases.SetupFile) *HabContext {
	logger := logger.NewLogger(startContext, "Hab")

	config := config.NewConfig(logger, defaultConfig)
	setup := setup.NewSetup(logger, config, defaultSetup)
	return &HabContext{
		startContext: startContext,
		log:          logger,
		config:       config,
		setup:        setup,
	}
}

func (h *HabContext) ParseCli(cmd *cli.Command) error {

	logLevel := cmd.String("loglevel")

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

	setup := cmd.String("setup")

	if setup == "" {
		h.setup.LoadDefaultSetup()
	} else {
		h.setup.LoadSetupFromYamlFile(setup)
	}

	return nil

}

func (h *HabContext) GetConfigValue(key string) string {
	return h.config.GetValue(key)
}

func (h *HabContext) SetConfigValue(key string, value string) {
	h.config.SetConfigValue(key, value)
}

func (h *HabContext) GetCurrentConfig() map[string]string {
	return h.config.GetCurrent()
}

func (h *HabContext) GetSetupContainers() []bases.SetupContainer {
	return h.setup.ContainersConfig
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
