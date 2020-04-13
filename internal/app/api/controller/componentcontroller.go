package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-ozzo/ozzo-routing/v2"

	"github.com/Kalinin-Andrey/rti-testing/pkg/errorshandler"
	"github.com/Kalinin-Andrey/rti-testing/pkg/log"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/component"
)

type componentController struct {
	Controller
	Service component.IService
	Logger  log.ILogger
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/component/
//	GET /api/component/{ID}
//	POST /api/component/
//	PUT /api/component/{ID}
//	DELETE /api/component/{ID}
func RegisterComponentHandlers(r *routing.RouteGroup, service component.IService, logger log.ILogger, authHandler routing.Handler) {
	c := componentController{
		Service:		service,
		Logger:			logger,
	}

	r.Get("/component", c.list)
	r.Get(`/component/<id:\d+>`, c.get)

	r.Use(authHandler)

	r.Post("/component", c.create)
	r.Put(`/component/<id:\d+>`, c.update)
	r.Delete(`/component/<id:\d+>`, c.delete)
}

// get method is for a getting a one enmtity by ID
func (c componentController) get(ctx *routing.Context) error {
	id, err := c.parseUint(ctx, "id")
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Info(errors.Wrapf(err, "Can not parse uint64 from %q", ctx.Param("id")))
		return errorshandler.BadRequest("id mast be a uint")
	}
	entity, err := c.Service.Get(ctx.Request.Context(), uint(id))
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(entity)
}

// list method is for a getting a list of entities
func (c componentController) list(ctx *routing.Context) error {

	items, err := c.Service.List(ctx.Request.Context())
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(items)
}

func (c componentController) create(ctx *routing.Context) error {
	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest("component invalid: " + err.Error())
	}

	if err := c.Service.Create(ctx.Request.Context(), entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.WriteWithStatus(entity, http.StatusCreated)
}

func (c componentController) update(ctx *routing.Context) error {
	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest("component invalid: " + err.Error())
	}

	if err := c.Service.Update(ctx.Request.Context(), entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.WriteWithStatus(entity, http.StatusCreated)
}


func (c componentController) delete(ctx *routing.Context) error {
	id, err := c.parseUint(ctx, "id")
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest("id must be uint")
	}

	if err := c.Service.Delete(ctx.Request.Context(), uint(id)); err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return errorshandler.Success()
}


