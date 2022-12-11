package handler

import (
	"fmt"
	"net/http"

	"digital-house/internal/dentist"
	"digital-house/pkg/util"
	"digital-house/pkg/web"

	"github.com/gin-gonic/gin"
)

type DentistController struct {
	service dentist.Service
}

func NewDentist(ctx *gin.Engine, service dentist.Service) {
	ec := &DentistController{service: service}

	re := ctx.Group("/api/v1/dentists")
	re.GET("/", ec.GetAll())
	re.GET("/:id", ec.GetByID())
	re.POST("/", ec.Create())
	re.PATCH("/:id", ec.Update())
	re.DELETE("/:id", ec.Delete())
}

func (c *DentistController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentists, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, dentists, ""))
	}
}

func (c *DentistController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := util.IDChecker(ctx)
		if err != nil {
			return
		}

		dentist, err := c.service.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, dentist, ""))
	}
}

func (c *DentistController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request dentist.RequestDentistCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		dentist, err := c.service.Create(request.FirstName, request.LastName, request.Registration)

		if err != nil {
			dentistExists := fmt.Sprintf("dentist with this registration %s exists", request.Registration)
			if err.Error() == dentistExists {
				ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
				return
			}
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, dentist, ""))

	}
}

func (c *DentistController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := util.IDChecker(ctx)
		if err != nil {
			return
		}

		var request dentist.RequestDentistUpdate

		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		dentist, err := c.service.Update(id, request.FirstName, request.LastName, request.Registration)

		if err != nil {
			dentistExists := fmt.Sprintf("dentist with this registration %s exists", request.Registration)
			if err.Error() == dentistExists {
				ctx.JSON(409, web.NewResponse(409, nil, err.Error()))
				return
			}
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, dentist, ""))
	}
}

func (c *DentistController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := util.IDChecker(ctx)
		if err != nil {
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}
