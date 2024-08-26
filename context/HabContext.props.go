package context

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/tuxounet/k-hab/bases"
)

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
	return h.setup.SetupContainers
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
	return h.workFolder
}

func (h *HabContext) SetLogLevel(level string) error {
	switch level {
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
	default:
		return errors.New("unknown log level")
	}
	return nil
}

func (h *HabContext) SetSetup(setup string) error {
	if setup == "" {
		return h.setup.LoadDefaultSetup()
	} else {
		return h.setup.LoadSetupFromYamlFile(setup)
	}

}
