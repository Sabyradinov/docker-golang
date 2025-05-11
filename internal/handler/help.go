package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Help struct {
}

func NewHelpHandler() *Help {
	return &Help{}
}

func (h Help) Get(ctx *gin.Context) {
	res := "test"

	ctx.JSON(http.StatusOK, res)
}
