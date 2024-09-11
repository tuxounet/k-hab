package context

import (
	"context"
	"log"
	"path"

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
	defaultConfig, err := utils.LoadYamlFromString[map[string]string](embedConfig.DefaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	config := config.NewConfig(defaultConfig)
	logFolder := config.GetValue("hab.logs.path")

	newContext := &HabContext{
		startContext: startContext,
		config:       config,
		workFolder:   workFolder,
	}
	storageRoot, err := newContext.GetStorageRoot()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(startContext, "Hab", path.Join(storageRoot, logFolder))
	config.SetLogger(logger)
	newContext.log = logger

	defaultSetup, err := utils.LoadYamlFromString[bases.SetupFile](embedConfig.DefaultSetup)
	if err != nil {
		log.Fatal(err)
	}

	setup := setup.NewSetup(logger, config, defaultSetup)

	newContext.setup = setup

	return newContext
}
