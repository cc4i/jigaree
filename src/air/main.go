// The purpose of this service is to simulated for quliafied data, and retrieve real data
// from AQI + Weathe API

package main

import (
	"air/aqi"
	"air/gen"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Configurue routes
func router(ctx context.Context, r *gin.Engine) *gin.Engine {

	r.GET("/ping", ping)
	r.GET("/metrics", metrics)
	r.GET("/readme", readme)
	r.GET("/aq/:city", aqByCity)
	r.GET("/aq/r/:num", aqWithRadom)

	return r
}

func aqWithRadom(c *gin.Context) {
	n, err := strconv.Atoi(c.Param("num"))
	log.Info().Int("num", n).Msg("ask for random data")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gen.RandomAQ(n))
	}

}

func aqByCity(c *gin.Context) {

	city := c.Param("city")
	bj, err := aqi.AirbyCity(c.Request.Context(), city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, bj)
	} else {
		c.JSON(http.StatusOK, bj)
	}

}

// Simple health check
func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// Prometheus metrics
func metrics(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

// Readme for air quality standard
func readme(c *gin.Context) {
	c.JSON(http.StatusOK, aqi.Readme())
}

var done = make(chan bool)
var ticker = time.NewTicker(5 * time.Second)

func repeat() {
	for {
		select {
		case <-done:
			fmt.Println("Stop reapted job.")
			ticker.Stop()
			return
		case <-ticker.C:
			// fmt.Println("Hello !!")
			resp, _ := http.Get("http://localhost:9011/ping")
			if resp.StatusCode == http.StatusOK {
				b, _ := ioutil.ReadAll(resp.Body)
				fmt.Printf("%s\n", b)
			}
		}
	}
}
func pHeaders() gin.HandlerFunc {
	//
	return func(c *gin.Context) {
		log.Info().Fields(c.Request.Header).Msg("request headers")

	}

}

// main
// default envrionments
// export REDIS_SERVER_ADDRESS="127.0.0.1:6379"
// export AQI_SERVER_URL="https://api.waqi.info"
// export AQI_SERVER_TOKEN="b0e78ca32d058a9170b6907c5214c0e946534cc9"
// export IP_STACK_SERVER_URL="http://api.ipstack.com"
// export IP_STACK_SERVER_TOKEN="ad7c6834f8dba51e8943d96d3742fcc5"
func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// go repeat()
	// To stop: done <- true

	gin.DisableConsoleColor()
	server := gin.Default()
	// server.Use(log.Logger_JSON())
	server.Use(pHeaders())
	log.Fatal().Err(router(context.Background(), server).Run("0.0.0.0:9011"))

}
