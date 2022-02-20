// The purpose of this service is to simulated for quliafied data, and retrieve real data
// from AQI + Weathe API

package main

import (
	"context"
	"maker/air"
	"maker/gen"
	log "maker/logging"
	"net/http"

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

	// Get AQIServer standard: Air Quality Index scale as defined by the US-EPA 2016 standard
	c.JSON(http.StatusOK, `
		{ 
			"version": "v1",
			{
				"Standard": "Air Quality Index scale as defined by the US-EPA 2016 standard.",
				"Definitions": [
					{
						"AQIServer": "0-50",
						"Level": "Good",
						"Implication": "Air quality is considered satisfactory, and air pollution poses little or no risk",
						"Caution": "None"
					},
					{
						"AQIServer": "51 -100",
						"Level": "Moderate",
						"Implication": "Air quality is acceptable; however, for some pollutants there may be a moderate health concern for a very small number of people who are unusually sensitive to air pollution.",
						"Caution": "Active children and adults, and people with respiratory disease, such as asthma, should limit prolonged outdoor exertion."
					},
					{
						"AQIServer": "101-150",
						"Level": "Unhealthy for Sensitive Groups",
						"Implication": "Members of sensitive groups may experience health effects. The general public is not likely to be affected.",
						"Caution": "Active children and adults, and people with respiratory disease, such as asthma, should limit prolonged outdoor exertion."
					},
					{
						"AQIServer": "151-200",
						"Level": "Unhealthy",
						"Implication": "Everyone may begin to experience health effects; members of sensitive groups may experience more serious health effects",
						"Caution": "Active children and adults, and people with respiratory disease, such as asthma, should avoid prolonged outdoor exertion; everyone else, especially children, should limit prolonged outdoor exertion"
					},
					{
						"AQIServer": "201-300",
						"Level": "Very Unhealthy",
						"Implication": "Health warnings of emergency conditions. The entire population is more likely to be affected.",
						"Caution": "Active children and adults, and people with respiratory disease, such as asthma, should avoid all outdoor exertion; everyone else, especially children, should limit outdoor exertion."
					},
					{
						"AQIServer": "300+",
						"Level": "Hazardous",
						"Implication": "Health alert: everyone may experience more serious health effects",
						"Caution": "Everyone should avoid all outdoor exertion"
					}
				]
			}
		}
	`)

}

// main
func main() {
	gin.DisableConsoleColor()
	server := gin.Default()
	server.Use(log.Logger_JSON())
	log.Lx.Fatal(router(context.Background(), server).Run("0.0.0.0:9011"))
}
