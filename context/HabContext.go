package context

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/logger"

	"github.com/tuxounet/k-hab/controllers/builder"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/controllers/ingress"

	"github.com/urfave/cli/v3"
)

type HabContext struct {
	startContext context.Context
	config       *config.Config
	log          *logger.Logger
	controllers  map[bases.HabControllers]bases.IController
}

func NewHabContext(startContext context.Context) *HabContext {

	return &HabContext{
		startContext: startContext,
		config:       config.NewConfig(),
		log:          logger.NewLogger(startContext),
	}
}

func (h *HabContext) ParseCli(cmd *cli.Command) error {
	quiet := cmd.Bool("quiet")
	println(quiet)

	logLevel := cmd.String("loglevel")
	println(logLevel)

	switch logLevel {
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

func (h *HabContext) Init() error {
	err := h.config.Load()
	if err != nil {
		return err
	}
	order := h.getControllersOrder()
	h.controllers = make(map[bases.HabControllers]bases.IController, len(order))

	for _, controllerKey := range order {
		var controller bases.IController

		switch controllerKey {
		case bases.DependenciesController:
			controller = dependencies.NewDependenciesController(h)

		case bases.BuilderController:
			controller = builder.NewBuilderController(h)

		case bases.IngressController:
			controller = ingress.NewHttpIngress(h)

		default:
			return errors.New("invalid controller name")
		}

		if controller == nil {
			return errors.New("iontroller is nil")
		}

		h.controllers[bases.HabControllers(controllerKey)] = controller
	}

	return nil
}

func (h *HabContext) GetHabConfig() bases.HabConfig {
	return h.config.HabConfig
}

func (h *HabContext) SetHabConfig(habConfig bases.HabConfig) {
	h.config.HabConfig = habConfig

}

func (h *HabContext) getControllersOrder() []bases.HabControllers {
	order := make([]bases.HabControllers, 0)
	order = append(order, bases.DependenciesController)
	order = append(order, bases.BuilderController)
	order = append(order, bases.IngressController)
	return order
}

func (h *HabContext) getReversedControllersOrder() []bases.HabControllers {
	order := make([]bases.HabControllers, 0)
	order = append(order, bases.IngressController)
	order = append(order, bases.BuilderController)
	order = append(order, bases.DependenciesController)
	return order
}

func (h *HabContext) GetController(controller bases.HabControllers) (bases.IController, error) {
	for key := range h.controllers {
		if key == bases.HabControllers(controller) {
			return h.controllers[key], nil
		}
	}
	return nil, errors.New("controller not found")

}

func (h *HabContext) Provision() error {
	for _, controllerKey := range h.getControllersOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Provision()
		if err != nil {
			return err
		}
	}
	h.log.InfoF("Hab Provisioned")
	return nil
}

func (h *HabContext) Start() error {

	//Ensure Provisioning
	err := h.Provision()
	if err != nil {
		return err
	}

	//Start
	for _, controllerKey := range h.getControllersOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Start()
		if err != nil {
			return err
		}
	}

	// ctx.Must(h.httpEgress.Start(ctx))

	// ctx.Must(h.upImages(ctx))
	// ctx.Must(h.upContainers(ctx))

	// ctx.Must(h.httpIngress.Start(ctx))
	// ctx.Must(h.startContainers(ctx))

	return nil

}

func (h *HabContext) Entry() error {

	//Ensure Start
	err := h.Start()

	if err != nil {
		return err
	}

	// container := h.getEntryContainer(ctx)
	// ctx.Must(container.waitReady(ctx))
	// ctx.Must(container.entry(ctx))

	err = h.Stop()
	if err != nil {
		return err
	}

	return nil

}

func (h *HabContext) Shell() error {
	// Ensure Start
	err := h.Start()

	if err != nil {
		return err
	}

	// container := h.getEntryContainer(ctx)
	// h.Must(container.shell(ctx))

	//Stop
	err = h.Stop()
	if err != nil {
		return err
	}

	return nil
}

func (h *HabContext) Stop() error {

	for _, controllerKey := range h.getReversedControllersOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}

		err = controller.Stop()
		if err != nil {
			return err
		}
	}
	return nil

	// ctx.Must(h.httpIngress.Stop(ctx))

	// 	ctx.Must(h.stopContainers(ctx))

	// 	ctx.Must(h.httpEgress.Stop(ctx))
	// })
}

func (h *HabContext) Rm() error {
	err := h.Stop()
	if err != nil {
		return err
	}
	for _, controllerKey := range h.getReversedControllersOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Rm()
		if err != nil {
			return err
		}
	}
	return nil

	// return h.ctx.Scope(h.scopeBase, "Rm", func(ctx *utils.ScopeContext) {
	// 	ctx.Must(h.Stop())
	// 	ctx.Must(h.downContainers(ctx))
	// 	ctx.Must(h.downImages(ctx))
	// 	ctx.Must(h.lxd.Down(ctx))
	// })
}

func (h *HabContext) Unprovision() error {

	err := h.Rm()
	if err != nil {
		return err
	}
	for _, controllerKey := range h.getReversedControllersOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Unprovision()
		if err != nil {
			return err
		}
	}
	return nil

	// 	lxdPresent := h.lxd.Present(ctx)
	// 	if lxdPresent {
	// 		ctx.Must(h.Rm())
	// 		ctx.Must(h.lxd.Unprovision(ctx))
	// 	}

	// 	builderPresent := h.builder.Present(ctx)
	// 	if builderPresent {
	// 		ctx.Must(h.builder.Unprovision(ctx))
	// 	}

	// })
}

func (h *HabContext) Nuke() error {
	err := h.Unprovision()
	if err != nil {
		return err
	}
	for _, controllerKey := range h.getReversedControllersOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Nuke()
		if err != nil {
			return err
		}
	}
	return nil

	// 	ctx.Must(h.Unprovision())
	// 	ctx.Must(h.nukeImages(ctx))
	// 	ctx.Must(h.lxd.Nuke(ctx))
	// 	ctx.Must(h.builder.Nuke(ctx))

}
