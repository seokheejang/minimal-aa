package start

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/seokheejang/minimal-aa/packages/bundler/logger"
)

func PrivateMode() {

	logr := logger.NewZeroLogr().
		WithName("minimal-aa").
		WithValues("bundler_mode", "private")

	logr.Info("Private Mode Start")

	// @TODO: eoa

	// @TODO: config

	// Init HTTP server
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/ping", func(g *gin.Context) {
		g.Status(http.StatusOK)
	})

	if err := r.Run(fmt.Sprintf(":%d", 4337)); err != nil {
		log.Fatal(err)
	}
}
