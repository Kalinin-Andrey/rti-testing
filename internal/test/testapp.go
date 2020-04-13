package test

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/offer"
	"github.com/Kalinin-Andrey/rti-testing/pkg/config"
	"github.com/Kalinin-Andrey/rti-testing/pkg/log"

	commonApp "github.com/Kalinin-Andrey/rti-testing/internal/app"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/user"
	"github.com/Kalinin-Andrey/rti-testing/internal/test/mock"
)

// New func is a constructor for the App
func NewCommonApp(cfg config.Configuration) *commonApp.App {
	logger, err := log.New(cfg.Log)
	if err != nil {
		panic(err)
	}

	app := &commonApp.App{
		Cfg:    cfg,
		Logger: logger,
		DB:     nil,
	}

	app.Domain.User.Repository = &mock.UserRepository{}
	app.Domain.User.Service = user.NewService(app.Domain.User.Repository, app.Logger)

	app.Domain.RuleApplicability.Repository = &mock.RuleApplicabilityRepository{}
	app.Domain.RuleApplicability.Service = ruleapplicability.NewService(app.Domain.RuleApplicability.Repository, app.Logger)

	app.Domain.Price.Repository = &mock.PriceRepository{}
	app.Domain.Price.Service = price.NewService(app.Domain.Price.Repository, app.Domain.RuleApplicability.Service, app.Logger)

	app.Domain.Component.Repository = &mock.ComponentRepository{}
	app.Domain.Component.Service = component.NewService(app.Domain.Component.Repository, app.Domain.Price.Service, app.Logger)

	app.Domain.Product.Repository = &mock.ProductRepository{}
	app.Domain.Product.Service = product.NewService(app.Domain.Product.Repository, app.Domain.Component.Service, app.Logger)

	app.Domain.Offer.Service = offer.NewService(app.Domain.Product.Service, app.Logger)

	app.SessionRepository = &mock.SessionRepository{}

	return app
}

