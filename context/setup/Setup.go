package setup

import (
	"os"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/utils"
)

type Setup struct {
	defaultSetup    bases.SetupFile
	config          *config.Config
	SetupContainers []bases.SetupContainer
	log             bases.ILogger
}

func NewSetup(logger bases.ILogger, config *config.Config, defaultSetup bases.SetupFile) *Setup {
	return &Setup{
		log:             logger.CreateSubLogger("Setup", logger),
		defaultSetup:    defaultSetup,
		config:          config,
		SetupContainers: make([]bases.SetupContainer, 0),
	}
}

func (s *Setup) LoadDefaultSetup() error {
	s.loadSetup(s.defaultSetup)
	return nil
}

func (s *Setup) LoadSetupFromYamlFile(file string) error {
	s.log.DebugF("Loading setup from file %s", file)
	setupBody, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	setupStr := string(setupBody)
	setupFile, err := utils.LoadYamlFromString[bases.SetupFile](setupStr)
	if err != nil {
		return err
	}
	s.loadSetup(setupFile)

	s.log.InfoF("Setup loaded from file %s", file)
	return nil

}

func (s *Setup) loadSetup(setup bases.SetupFile) {

	if setup.Config != nil {
		for key, value := range setup.Config {
			s.config.SetConfigValue(key, value)
		}
	}

	s.SetupContainers = append(s.SetupContainers, setup.Containers...)

}
