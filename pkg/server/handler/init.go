package handler

import "github.com/gin-gonic/gin"

type handler struct {
	customer  customerHandlerInterface
	websocket websocketHandlerInterface
}

var h handler

func init() {
	h = handler{
		customer:  getCustomerHandler(),
		websocket: getWebsocketHandler(),
	}
}
func Setup(app *gin.Engine) {
	h.customer.setupCustomerHandler(app)
	h.websocket.setupWebsocketHandler(app)
}
