package main

import (
	"context"
	"net/http"

	docs "gate/docs"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func pHeaders() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Info().Fields(c.Request.Header).Msg("request headers")

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

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	gin.DisableConsoleColor()
	server := gin.Default()
	server.Use(pHeaders())
	log.Fatal().Err(router(context.Background(), server).Run("0.0.0.0:9000"))
}
