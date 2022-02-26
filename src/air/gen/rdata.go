package gen

import (
	"air/aqi"
	"air/log"
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

func cities(n int) [][]string {

	var cs [][]string
	for i := 0; i < n; i++ {
		cs = append(cs, []string{
			fmt.Sprintf("Tekawa%d", i),
			fmt.Sprintf("特卡瓦%d", i),
		})
	}
	return cs
}

func RandomAQ(n int) []aqi.AirQuality {
	var rAQ []aqi.AirQuality
	cs := cities(n)
	for _, c := range cs {
		originAirQuality := aqi.OriginAirQuality{
			Status: "ok",
			Data: aqi.OriginData{
				AQI:          63,
				StationIndex: rand.Intn(9999),
				City: aqi.OriginCity{
					Geo: []float64{
						39.954592,
						116.468117,
					},
					Name: fmt.Sprintf("%s (%s)", c[0], c[1]),
				},
				IAQI: aqi.OriginIAQI{
					Co: aqi.OValue{
						V: rand.Float64(),
					},
					H: aqi.OValue{
						V: rand.Float64(),
					},
					No2: aqi.OValue{
						V: rand.Float64(),
					},
					O3: aqi.OValue{
						V: rand.Float64(),
					},
					P: aqi.OValue{
						V: rand.Float64(),
					},
					Pm10: aqi.OValue{
						V: rand.Float64(),
					},
					Pm25: aqi.OValue{
						V: rand.Float64(),
					},
					So2: aqi.OValue{
						V: rand.Float64(),
					},
					T: aqi.OValue{
						V: rand.Float64(),
					},
					W: aqi.OValue{
						V: 3.6,
					},
				},
				OriginTime: aqi.OriginTime{
					S:  "2099-09-09 09:00:00",
					TZ: "+08:00",
					V:  int(time.Now().Unix()),
				},
			},
		}
		log.Lx.WithFields(logrus.Fields{
			"air": originAirQuality,
		}).Info("original air quality data")
		oa := aqi.Copy2AirQuality(originAirQuality)
		log.Lx.WithFields(logrus.Fields{
			"new_air": oa,
		}).Info("coverted air quality data")
		rAQ = append(rAQ, oa)
	}

	return rAQ
}
