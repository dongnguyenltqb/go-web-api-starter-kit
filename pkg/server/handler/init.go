package handler

import "github.com/gin-gonic/gin"

type handler struct {
	customer customerHandlerInterface
}

var h handler

func init() {
	h = handler{
		customer: getCustomerHandler(),
	}
}
func Setup(app *gin.Engine) {
	h.customer.setupCustomerHandler(app)
}
