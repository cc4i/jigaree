package main

import (
	"context"
	"gate/log"
	"net/http"

	docs "gate/docs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func pHeaders() gin.HandlerFunc {
	//
	return func(c *gin.Context) {
		log.Lx.WithFields(logrus.Fields{
			"headers": c.Request.Header,
		}).Info("request headers")
	}

}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func router(ctx context.Context, r *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
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
