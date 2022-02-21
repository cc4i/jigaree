// The purpose of this service is to simulated for quliafied data, and retrieve real data
// from AQI + Weathe API

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"maker/air"
	"maker/gen"
	"maker/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Configurue routes
func router(ctx context.Context, r *gin.Engine) *gin.Engine {

	r.GET("/ping", ping)
	r.GET("/metrics", metrics)
	r.GET("/readme", readme)
	r.GET("/aq/:city", aqByCity)

	return r
}

func aqByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "_r" {
		c.JSON(http.StatusOK, gen.RandomAQ())
		return
	}
	bj, err := air.AirbyCity(c.Request.Context(), city)
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
	c.JSON(http.StatusOK, air.Readme())
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

// main
func main() {

	// go repeat()
	// To stop: done <- true

	gin.DisableConsoleColor()
	server := gin.Default()
	server.Use(log.Logger_JSON())
	log.Lx.Fatal(router(context.Background(), server).Run("0.0.0.0:9011"))

}
