package handler

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type websocketHandlerInterface interface {
	setupConn(c *gin.Context)
	handleConn(conn *websocket.Conn)
	setupWebsocketHandler(app *gin.Engine)
}

type websocketHandler struct {
}

func getWebsocketHandler() websocketHandlerInterface {
	return &websocketHandler{}
}
func (h *websocketHandler) setupWebsocketHandler(app *gin.Engine) {
	app.Any("/websocket", h.setupConn)
}

func (h *websocketHandler) setupConn(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	h.handleConn(conn)
}

func (h *websocketHandler) writeMessage(conn *websocket.Conn, messageType int, b []byte) error {
	w, err := conn.NextWriter(messageType)
	if err != nil {
		return err
	}
	defer w.Close()
	if _, err = w.Write(b); err != nil {
		return err
	}
	return nil
}

func (h *websocketHandler) handleConn(conn *websocket.Conn) {
	go func(conn *websocket.Conn) {
		for {
			<-time.After(2 * time.Second)
			if err := h.writeMessage(conn, websocket.TextMessage, []byte(fmt.Sprintf("Bây giờ là : %v \r\n", time.Now().UTC().String()))); err != nil {
				return
			}
		}
	}(conn)

	for {
		messageType, r, err := conn.NextReader()
		logrus.Info(messageType)
		if err != nil {
			logrus.Error(err)
			return
		}
		p, err := ioutil.ReadAll(r)
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Info(string(p))
		}

	}
}
