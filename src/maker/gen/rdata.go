package gen

import (
	"fmt"
	"maker/air"
	"maker/log"
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

func RandomAQ(n int) []air.AirQuality {
	var rAQ []air.AirQuality
	cs := cities(n)
	for _, c := range cs {
		originAirQuality := air.OriginAirQuality{
			Status: "ok",
			Data: air.OriginData{
				AQI:          63,
				StationIndex: rand.Intn(9999),
				City: air.OriginCity{
					Geo: []float64{
						39.954592,
						116.468117,
					},
					Name: fmt.Sprintf("%s (%s)", c[0], c[1]),
				},
				IAQI: air.OriginIAQI{
					Co: air.OValue{
						V: rand.Float64(),
					},
					H: air.OValue{
						V: rand.Float64(),
					},
					No2: air.OValue{
						V: rand.Float64(),
					},
					O3: air.OValue{
						V: rand.Float64(),
					},
					P: air.OValue{
						V: rand.Float64(),
					},
					Pm10: air.OValue{
						V: rand.Float64(),
					},
					Pm25: air.OValue{
						V: rand.Float64(),
					},
					So2: air.OValue{
						V: rand.Float64(),
					},
					T: air.OValue{
						V: rand.Float64(),
					},
					W: air.OValue{
						V: 3.6,
					},
				},
				OriginTime: air.OriginTime{
					S:  "2099-09-09 09:00:00",
					TZ: "+08:00",
					V:  int(time.Now().Unix()),
				},
			},
		}
		log.Lx.WithFields(logrus.Fields{
			"air": originAirQuality,
		}).Info("original air quality data")
		oa := air.Copy2AirQuality(originAirQuality)
		log.Lx.WithFields(logrus.Fields{
			"new_air": oa,
		}).Info("coverted air quality data")
		rAQ = append(rAQ, oa)
	}

	return rAQ
}
