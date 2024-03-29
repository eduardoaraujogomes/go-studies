package handler

import (
	"fmt"
	"net/http"

	inboundOrders "digital-house/internal/inbound_orders"
	"digital-house/pkg/util"
	"digital-house/pkg/web"

	"github.com/gin-gonic/gin"
)

type InboundOrdersController struct {
	service inboundOrders.Service
}

func NewInboundOrders(ctx *gin.Engine, service inboundOrders.Service) {
	ioc := &InboundOrdersController{service: service}

	ior := ctx.Group("/api/v1")
	{
		ior.GET("/dentists/reportInboundOrders", ioc.ReportInboundOrders())
		ior.POST("/inboundOrders", ioc.Create())
	}
}

func (c *InboundOrdersController) ReportInboundOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentistID := ctx.Query("id")
		if dentistID != "" {
			id, err := util.IDChecker(ctx)
			if err != nil {
				return
			}
			reportInboundOrder, err := c.service.GetByID(id)

			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reportInboundOrder, ""))
		} else {
			reportInboundOrders, err := c.service.GetAll()
			if err != nil {
				ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
				return
			}
			ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, reportInboundOrders, ""))
		}
	}
}

func (c *InboundOrdersController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request inboundOrders.InboundOrdersCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		inboundOrders, err := c.service.Create(request)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, inboundOrders, ""))
	}
}
