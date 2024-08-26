package context

import (
	"context"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/logger"
	"github.com/tuxounet/k-hab/context/setup"
)

type HabContext struct {
	startContext context.Context
	workFolder   string
	config       *config.Config
	setup        *setup.Setup
	log          *logger.Logger
	controllers  map[bases.HabControllers]bases.IController
}

func NewHabContext(startContext context.Context,
	defaultConfig map[string]string,
	defaultSetup bases.SetupFile,
	workFolder string) *HabContext {
	logger := logger.NewLogger(startContext, "Hab")

	config := config.NewConfig(logger, defaultConfig)
	setup := setup.NewSetup(logger, config, defaultSetup)
	return &HabContext{
		startContext: startContext,
		log:          logger,
		config:       config,
		setup:        setup,
		workFolder:   workFolder,
	}
}
