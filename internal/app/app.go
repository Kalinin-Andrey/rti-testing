package app

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/offer"
	golog "log"

	"github.com/Kalinin-Andrey/rti-testing/pkg/config"
	"github.com/Kalinin-Andrey/rti-testing/pkg/log"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/db"
	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/session"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
	//"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	//"github.com/Kalinin-Andrey/rti-testing/internal/domain/offer"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/price"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/ruleapplicability"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/user"
	dbrep "github.com/Kalinin-Andrey/rti-testing/internal/infrastructure/repository/db"
)

type IApp interface {
	// Run is func to run the App
	Run() error
}

// App struct is the common part of all applications
type App struct {
	Cfg					config.Configuration
	Logger				log.ILogger
	DB					db.IDB
	Domain				Domain
	SessionRepository	session.IRepository
}

// Domain is a Domain Layer Entry Point
type Domain struct {
	User struct {
		Repository		user.IRepository
		Service			user.IService
	}
	Product struct {
		Repository		product.IRepository
		Service			product.IService
	}
	Component struct {
		Repository		component.IRepository
		Service			component.IService
	}
	Price struct {
		Repository		price.IRepository
		Service			price.IService
	}
	RuleApplicability struct {
		Repository		ruleapplicability.IRepository
		Service			ruleapplicability.IService
	}
	Offer struct {
		Service			offer.IService
	}
}

const (
	EntityNameComponent			= "component"
	EntityNamePrice				= "price"
	EntityNameProduct			= "product"
	EntityNameRuleApplicability	= "ruleapplicability"
	EntityNameUser				= "user"
)

// New func is a constructor for the App
func New(cfg config.Configuration) *App {
	logger, err := log.New(cfg.Log)
	if err != nil {
		panic(err)
	}

	db, err := db.New(cfg.DB, logger)
	if err != nil {
		panic(err)
	}

	app := &App{
		Cfg:    cfg,
		Logger: logger,
		DB:     db,
	}
	var ok bool

	app.Domain.User.Repository, ok	= app.getDBRepo(EntityNameUser).(user.IRepository)
	if !ok {
		golog.Fatalf("Can not cast DB repository for entity %q to %v.IRepository. Repo: %v", EntityNameUser, EntityNameUser, app.getDBRepo(EntityNameUser))
	}
	app.Domain.User.Service = user.NewService(app.Domain.User.Repository, app.Logger)

	app.Domain.RuleApplicability.Repository, ok	= app.getDBRepo(EntityNameRuleApplicability).(ruleapplicability.IRepository)
	if !ok {
		golog.Fatalf("Can not cast DB repository for entity %q to %v.IRepository. Repo: %v", EntityNameRuleApplicability, EntityNameRuleApplicability, app.getDBRepo(EntityNameRuleApplicability))
	}
	app.Domain.RuleApplicability.Service = ruleapplicability.NewService(app.Domain.RuleApplicability.Repository, app.Logger)

	app.Domain.Price.Repository, ok	= app.getDBRepo(EntityNamePrice).(price.IRepository)
	if !ok {
		golog.Fatalf("Can not cast DB repository for entity %q to %v.IRepository. Repo: %v", EntityNamePrice, EntityNamePrice, app.getDBRepo(EntityNamePrice))
	}
	app.Domain.Price.Service = price.NewService(app.Domain.Price.Repository, app.Domain.RuleApplicability.Service, app.Logger)

	app.Domain.Component.Repository, ok	= app.getDBRepo(EntityNameComponent).(component.IRepository)
	if !ok {
		golog.Fatalf("Can not cast DB repository for entity %q to %v.IRepository. Repo: %v", EntityNameComponent, EntityNameComponent, app.getDBRepo(EntityNameComponent))
	}
	app.Domain.Component.Service = component.NewService(app.Domain.Component.Repository, app.Domain.Price.Service, app.Logger)

	app.Domain.Product.Repository, ok	= app.getDBRepo(EntityNameProduct).(product.IRepository)
	if !ok {
		golog.Fatalf("Can not cast DB repository for entity %q to %v.IRepository. Repo: %v", EntityNameProduct, EntityNameProduct, app.getDBRepo(EntityNameProduct))
	}
	app.Domain.Product.Service = product.NewService(app.Domain.Product.Repository, app.Domain.Component.Service, app.Logger)

	app.Domain.Offer.Service = offer.NewService(app.Domain.Product.Service, app.Logger)

	if app.SessionRepository, err = dbrep.NewSessionRepository(app.DB, app.Logger, app.Domain.User.Repository); err != nil {
		golog.Fatalf("Can not get new SessionRepository err: %v", err)
	}

	return app
}

// Run is func to run the App
func (app *App) Run() error {
	return nil
}

func (app *App) getDBRepo(entityName string) (repo dbrep.IRepository) {
	var err error

	if repo, err = dbrep.GetRepository(app.DB, app.Logger, entityName); err != nil {
		golog.Fatalf("Can not get db repository for entity %q, error happened: %v", entityName, err)
	}
	return repo
}
