package web

import (
	"eastv2/define"
	"eastv2/pb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tnnmigga/corev2/log"
	"github.com/tnnmigga/corev2/message"
)

func (m *web) initHandle() {
	m.GET("/index", index)
	m.POST("/login", loginOrRegister)
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func loginOrRegister(c *gin.Context) {
	data := struct {
		Account  string
		Password string
	}{}
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	resp, err := message.RequestAny[pb.AuthOrCreateAccountResp](define.SERV_MASTER, &pb.AuthOrCreateAccountReq{
		Account:  data.Account,
		Password: data.Password,
	})
	if err != nil {
		log.Error(err)
		c.String(http.StatusInternalServerError, "")
	}
	c.JSON(http.StatusOK, resp)
}
