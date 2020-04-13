package controller

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-ozzo/ozzo-routing/v2"

	"github.com/Kalinin-Andrey/rti-testing/pkg/errorshandler"
	"github.com/Kalinin-Andrey/rti-testing/pkg/log"

	"github.com/Kalinin-Andrey/rti-testing/internal/pkg/apperror"

	"github.com/Kalinin-Andrey/rti-testing/internal/domain/product"
)

type productController struct {
	Controller
	Service product.IService
	Logger  log.ILogger
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/product/
//	GET /api/product/{ID}
//	POST /api/product/
//	PUT /api/product/{ID}
//	DELETE /api/product/{ID}
func RegisterProductHandlers(r *routing.RouteGroup, service product.IService, logger log.ILogger, authHandler routing.Handler) {
	c := productController{
		Service:		service,
		Logger:			logger,
	}

	r.Get("/product", c.list)
	r.Get(`/product/<id:\d+>`, c.get)

	r.Use(authHandler)

	r.Post("/product", c.create)
	r.Put(`/product/<id:\d+>`, c.update)
	r.Delete(`/product/<id:\d+>`, c.delete)
}

// get method is for a getting a one enmtity by ID
func (c productController) get(ctx *routing.Context) error {
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
func (c productController) list(ctx *routing.Context) error {

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

func (c productController) create(ctx *routing.Context) error {
	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest("product invalid: " + err.Error())
	}

	if err := c.Service.Create(ctx.Request.Context(), entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.WriteWithStatus(entity, http.StatusCreated)
}

func (c productController) update(ctx *routing.Context) error {
	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest("product invalid: " + err.Error())
	}

	if err := c.Service.Update(ctx.Request.Context(), entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.WriteWithStatus(entity, http.StatusCreated)
}


func (c productController) delete(ctx *routing.Context) error {
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


