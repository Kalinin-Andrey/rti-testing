package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/go-ozzo/ozzo-routing/v2/slash"

	"github.com/Kalinin-Andrey/rti-testing/pkg/accesslog"
	"github.com/Kalinin-Andrey/rti-testing/pkg/config"
	"github.com/Kalinin-Andrey/rti-testing/pkg/errorshandler"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/auth"

	commonApp "github.com/Kalinin-Andrey/rti-testing/internal/app"
	"github.com/Kalinin-Andrey/rti-testing/internal/app/api/controller"
)

// Version of API
const Version = "1.0.0"

// App is the application for API
type App struct {
	*commonApp.App
	Server		*http.Server
}

// New func is a constructor for the ApiApp
func New(commonApp *commonApp.App, cfg config.Configuration) *App {
	app := &App{
		App:	commonApp,
		Server:	nil,
	}

	// build HTTP server
	server := &http.Server{
		Addr:		cfg.Server.HTTPListen,
		Handler:	app.buildHandler(),
	}
	app.Server = server

	return app
}

func (app *App) buildHandler() *routing.Router {
	router := routing.New()

	router.Use(
		accesslog.Handler(app.Logger),
		slash.Remover(http.StatusMovedPermanently),
		errorshandler.Handler(app.Logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)
	//router.NotFound(file.Content("website/index.html"))

	// serve index file
	/*router.Get("/", file.Content("website/index.html"))
	router.Get("/favicon.ico", file.Content("website/favicon.ico"))

	// serve files under the "static" subdirectory
	router.Get("/static/*", file.Server(file.PathMap{
		"/static/": "/website/static/",
	}))*/

	rg := router.Group("/api")

	authHandler := auth.Handler(app.Cfg.JWTSigningKey, app.DB, app.Logger, app.SessionRepository)

	auth.RegisterHandlers(rg.Group(""),
		auth.NewService(app.Cfg.JWTSigningKey, app.Cfg.JWTExpiration, app.Domain.User.Service, app.Logger),
		app.Logger,
	)

	app.RegisterHandlers(rg, authHandler)

	return 	router
}

// Run is func to run the ApiApp
func (app *App) Run() error {
	go func() {
		defer func() {
			if err := app.DB.DB().Close(); err != nil {
				app.Logger.Error(err)
			}

			err := app.Logger.Sync()
			if err != nil {
				log.Println(err.Error())
			}
		}()
		// start the HTTP server with graceful shutdown
		routing.GracefulShutdown(app.Server, 10*time.Second, app.Logger.Infof)
	}()
	app.Logger.Infof("server %v is running at %v", Version, app.Server.Addr)
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func (app *App) RegisterHandlers(rg *routing.RouteGroup, authHandler routing.Handler) {

	//controller.RegisterUserHandlers(rg, app.Domain.User.Service, app.Logger, authHandler)
	controller.RegisterProductHandlers(rg, app.Domain.Product.Service, app.Logger, authHandler)
	controller.RegisterComponentHandlers(rg, app.Domain.Component.Service, app.Logger, authHandler)
	controller.RegisterPriceHandlers(rg, app.Domain.Price.Service, app.Logger, authHandler)
	controller.RegisterRuleApplicabilityHandlers(rg, app.Domain.RuleApplicability.Service, app.Logger, authHandler)
	controller.RegisterOfferHandlers(rg, app.Domain.Offer.Service, app.Logger, authHandler)

}
