package handler

import (
	"ganja/pkg/server/usecase"

	"github.com/gin-gonic/gin"
)

type customerHandlerInterface interface {
	getById(c *gin.Context)
	setupCustomerHandler(app *gin.Engine)
}

type customerHandler struct {
	customerUsecase usecase.CustomerUsecaseInterface
}

func getCustomerHandler() customerHandlerInterface {
	return &customerHandler{
		customerUsecase: usecase.GetCustomerUsecase(),
	}
}

func (h *customerHandler) setupCustomerHandler(app *gin.Engine) {
	customerGroup := app.Group("/customers")
	{
		customerGroup.GET("/:id", h.getById)
	}
}

func (h *customerHandler) getById(c *gin.Context) {
	id := c.Param("id")
	cus, err := h.customerUsecase.GetById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"ok":   true,
		"data": cus,
	})
}
