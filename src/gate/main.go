package main

import (
	"context"
	"gate/log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func pHeaders() gin.HandlerFunc {
	//
	return func(c *gin.Context) {
		log.Lx.WithFields(logrus.Fields{
			"headers": c.Request.Header,
		}).Info("request headers")
	}

}
func router(ctx context.Context, r *gin.Engine) *gin.Engine {

	return r
}

func main() {

	// go repeat()
	// To stop: done <- true

	gin.DisableConsoleColor()
	server := gin.Default()
	server.Use(log.Logger_JSON())
	server.Use(pHeaders())
	log.Lx.Fatal(router(context.Background(), server).Run("0.0.0.0:9000"))
}
