package controller

import (
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/condition"
	"github.com/Kalinin-Andrey/rti-testing/internal/domain/offer"
	"net/http"

	"github.com/pkg/errors"

	"github.com/go-ozzo/ozzo-routing/v2"

	"github.com/Kalinin-Andrey/rti-testing/pkg/errorshandler"
	"github.com/Kalinin-Andrey/rti-testing/pkg/log"
)

type offerController struct {
	Controller
	Service offer.IService
	Logger  log.ILogger
}

const usualConditions = 2

// RegisterHandlers sets up the routing of the HTTP handlers.
//	POST /api/product/{ID}/offer
func RegisterOfferHandlers(r *routing.RouteGroup, service offer.IService, logger log.ILogger, authHandler routing.Handler) {
	c := offerController{
		Service:		service,
		Logger:			logger,
	}

	r.Post(`/product/<productId:\d+>/offer`, c.calculate)

	//r.Use(authHandler)
}


func (c offerController) calculate(ctx *routing.Context) error {
	productId, err := c.parseUint(ctx, "productId")
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Info(errors.Wrapf(err, "Can not parse uint64 from %q", ctx.Param("id")))
		return errorshandler.BadRequest("id mast be a uint")
	}

	conditions := make([]condition.Condition, 0, usualConditions)
	if err := ctx.Read(&conditions); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	for _, c := range conditions {
		if err := c.Validate(); err != nil {
			return errorshandler.BadRequest("condition invalid: " + err.Error())
		}
	}

	offer, err := c.Service.CalculateByProductID(ctx.Request.Context(), productId, conditions)
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError(err.Error())
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.WriteWithStatus(offer, http.StatusOK)
}



