package service

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	pb "weather/protos"
)

var (
	WeatherServer      string
	WeatherServerToken string
	localCaches        = make(map[string][]byte)
)

type WeatheService struct {
	pb.UnimplementedWeatherServiceServer
}

// https://api.weatherapi.com/ - chuancc+weatherapi@gmail.com
func init() {
	WeatherServer = os.Getenv("WEATHER_SERVER_URL")
	WeatherServerToken = os.Getenv("WEATHER_SERVER_TOKEN")
	if WeatherServer == "" || WeatherServerToken == "" {
		log.Printf("WeatherServer => %s, WeatherServerToken => %s\n", WeatherServer, WeatherServerToken)
		log.Fatal("Failed to initail environments setting, pls configure & Reboot.")
	}
}

func (s *WeatheService) WeatherbyCity(ctx context.Context, in *pb.CityWeatherRequest) (*pb.CityWeatherResponse, error) {

	// Call API if [(Localtime + 600) < time.Now().Unix()]
	if localCaches[in.City] != nil {
		old := buildCityWeatherResponse(localCaches[in.City])
		if (old.Localtime + 600) > time.Now().Unix() {
			log.Printf("Using cached data for %s\n", in.City)
			return old, nil
		}
	}

	url := fmt.Sprintf("%s/v1/current.json?key=%s&q=%s&aqi=no", WeatherServer, WeatherServerToken, in.City)
	log.Printf("Call url => %s\n", url)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Request %s with error => %s\n", url, err.Error())
		return &pb.CityWeatherResponse{}, err
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	localCaches[in.City] = buf
	return buildCityWeatherResponse(buf), nil
}

func buildCityWeatherResponse(buf []byte) *pb.CityWeatherResponse {

	var data map[string]interface{}
	json.Unmarshal(buf, &data)
	loc := data["location"].(map[string]interface{})
	current := data["current"].(map[string]interface{})
	condition := current["condition"].(map[string]interface{})
	log.Println(loc["name"], time.Now().Unix())

	return &pb.CityWeatherResponse{
		City:          loc["name"].(string),
		Region:        loc["region"].(string),
		Country:       loc["country"].(string),
		Tz:            loc["tz_id"].(string),
		Localtime:     int64(loc["localtime_epoch"].(float64)),
		LastUpdate:    int64(current["last_updated_epoch"].(float64)),
		ConditionText: condition["text"].(string),
		ConditionIcon: condition["icon"].(string),
		WindKph:       current["wind_kph"].(float64),
		WindDegree:    current["wind_degree"].(float64),
		WindDir:       current["wind_dir"].(string),
		PressureMb:    current["pressure_mb"].(float64),
		PressureIn:    current["pressure_in"].(float64),
		PrecipMm:      current["precip_mm"].(float64),
		PrecipIn:      current["precip_in"].(float64),
		Humidity:      current["humidity"].(float64),
		Cloud:         current["cloud"].(float64),
		FeelslikeC:    current["feelslike_c"].(float64),
		VisKm:         current["vis_km"].(float64),
		Uv:            current["uv"].(float64),
		GustKph:       current["gust_kph"].(float64),
	}

}
