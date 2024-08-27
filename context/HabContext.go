package context

import (
	"context"
	"log"

	"github.com/tuxounet/k-hab/bases"
	embedConfig "github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/logger"
	"github.com/tuxounet/k-hab/context/setup"
	"github.com/tuxounet/k-hab/utils"
)

type HabContext struct {
	startContext context.Context
	workFolder   string
	config       *config.Config
	setup        *setup.Setup
	log          *logger.Logger
	controllers  map[bases.HabControllers]bases.IController
}

func NewHabContext(startContext context.Context, workFolder string) *HabContext {
	logger := logger.NewLogger(startContext, "Hab")

	defaultConfig, err := utils.LoadYamlFromString[map[string]string](embedConfig.DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	defaultSetup, err := utils.LoadYamlFromString[bases.SetupFile](embedConfig.DefaultSetup)
	if err != nil {
		log.Fatal(err)
	}

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
