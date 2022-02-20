package gen

import (
	"fmt"
	"maker/air"
	log "maker/logging"

	"github.com/sirupsen/logrus"
)

func cities() [][]string {

	var cs [][]string
	for i := 0; i < 50; i++ {
		cs = append(cs, []string{
			fmt.Sprintf("Tekawa%d", i),
			fmt.Sprintf("特卡瓦%d", i),
		})
	}
	return cs
}

func RandomAQ() []air.AirQuality {
	var rAQ []air.AirQuality
	cs := cities()
	for _, c := range cs {
		originAirQuality := air.OriginAirQuality{
			Status: "ok",
			Data: air.OriginData{
				AQI:          63,
				StationIndex: 1451,
				City: air.OriginCity{
					Geo: []float64{
						39.954592,
						116.468117,
					},
					Name: fmt.Sprintf("%s (%s)", c[0], c[1]),
				},
				IAQI: air.OriginIAQI{
					Co: air.OValue{
						V: 4.6,
					},
					H: air.OValue{
						V: 19,
					},
					No2: air.OValue{
						V: 5.5,
					},
					O3: air.OValue{
						V: 37.8,
					},
					P: air.OValue{
						V: 1020,
					},
					Pm10: air.OValue{
						V: 56,
					},
					Pm25: air.OValue{
						V: 63,
					},
					So2: air.OValue{
						V: 3.6,
					},
					T: air.OValue{
						V: 15,
					},
					W: air.OValue{
						V: 3.6,
					},
				},
				OriginTime: air.OriginTime{
					S:  "2022-10-01 17:00:00",
					TZ: "+08:00",
					V:  1586365200,
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
