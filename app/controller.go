package app

import (
	"github.com/gin-gonic/gin"
	"github.com/step_ddd/db"
	"github.com/step_ddd/domain/parking"
	"github.com/step_ddd/event"
	"net/http"
	"strings"
)

var (
	ci = &parking.CheckInCommandHandler{Repository: db.DB}
)

func registerHandlerFunc(r *gin.Engine) {
	r.GET("/checkIn", handleCheckIn)
}

func handleCheckIn(c *gin.Context) {
	plate := c.Query("plate")

	if err := ci.Handle(event.MQ, &parking.CheckInCommand{Plate: plate}); err != nil {
		fail(c, err.Error())
		return
	}
	ok(c)
}

func ok(c *gin.Context, msg ...string) {
	c.JSON(http.StatusOK, map[string]string{
		"result": "success",
		"msg":    strings.Join(msg, ","),
	})
}

func fail(c *gin.Context, msg ...string) {
	c.JSON(http.StatusOK, map[string]string{
		"result": "fail",
		"msg":    strings.Join(msg, ","),
	})
}
