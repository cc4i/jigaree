package air

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"maker/log"

	"github.com/sirupsen/logrus"
)

var (
	Client             HttpClient
	AQIServer          string // "https://api.waqi.info"
	AQIServerToken     string // "b0e78ca32d058a9170b6907c5214c0e946534cc9"
	IpStackServer      string // "http://api.ipstack.com"
	IpStackServerToken string // "ad7c6834f8dba51e8943d96d3742fcc5"

	//api.ipstack.com/127.0.0.1?access_key=ad7c6834f8dba51e8943d96d3742fcc5
	//https://ipapi.co/json
	//https://ipapi.co/8.8.8.8/json
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type IpStack struct {
	Ip            string  `json:"ip"`
	Type          string  `json:"type"`
	ContinentCode string  `json:"continent_code"`
	ContinentName string  `json:"continent_name"`
	CountryCode   string  `json:"country_code"`
	CountryName   string  `json:"country_name"`
	RegionCode    string  `json:"region_code"`
	RegionName    string  `json:"region_name"`
	City          string  `json:"city"`
	Cip           string  `json:"zip"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type ApiError struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

type AirQuality struct {
	IndexCityVHash string `json:"index_city_v_hash"`
	IndexCity      string `json:"index_city"`
	StationIndex   int    `json:"idx"`
	AQI            int    `json:"aqi"`
	City           string `json:"city"`
	CityCN         string `json:"city_cn"`
	Latitude       string `json:"lat"`
	Longitude      string `json:"lng"`
	Co             string `json:"co"`
	H              string `json:"h"`
	No2            string `json:"no2"`
	O3             string `json:"o3"`
	P              string `json:"p"`
	Pm10           string `json:"pm10"`
	Pm25           string `json:"pm25"`
	So2            string `json:"so2"`
	T              string `json:"t"`
	W              string `json:"w"`
	S              string `json:"s"`  //Local measurement time
	TZ             string `json:"tz"` //Station timezone
	V              int    `json:"v"`
}

type OriginAirQuality struct {
	Status string     `json:"status"`
	Data   OriginData `json:"data"`
}

type OriginData struct {
	AQI          int        `json:"aqi"`
	StationIndex int        `json:"idx"`
	City         OriginCity `json:"city"`
	IAQI         OriginIAQI `json:"iaqi"`
	OriginTime   OriginTime `json:"time"`
}

type OriginCity struct {
	Geo  []float64 `json:"geo"`
	Name string    `json:"name"`
}

type OriginIAQI struct {
	Co   OValue `json:"co"`
	H    OValue `json:"h"`
	No2  OValue `json:"no2"`
	O3   OValue `json:"o3"`
	P    OValue `json:"p"`
	Pm10 OValue `json:"pm10"`
	Pm25 OValue `json:"pm25"`
	So2  OValue `json:"so2"`
	T    OValue `json:"t"`
	W    OValue `json:"w"`
}

type OValue struct {
	V float64 `json:"v"`
}

type OriginTime struct {
	S  string `json:"s"`  //Local measurement time
	TZ string `json:"tz"` //Station timezone
	V  int    `json:"v"`
}

// Initial function for AQI APIs
func init() {
	// Initial ENVs
	AQIServer = os.Getenv("AQI_SERVER_URL")
	AQIServerToken = os.Getenv("AQI_SERVER_TOKEN")
	IpStackServer = os.Getenv("IP_STACK_SERVER_URL")
	IpStackServerToken = os.Getenv("IP_STACK_SERVER_TOKEN")
	if AQIServer == "" || IpStackServer == "" || AQIServerToken == "" || IpStackServerToken == "" {
		log.Lx.Fatal("Retrieving servers' address were failed. Check out environments setting & Reboot.")
		log.Lx.WithFields(logrus.Fields{
			"AQI_SERVER_URL":        AQIServer,
			"AQI_SERVER_TOKEN":      AQIServerToken,
			"IP_STACK_SERVER_URL":   IpStackServer,
			"IP_STACK_SERVER_TOKEN": IpStackServer,
		}).Error("Failed to initail environments setting, pls configure & Reboot.")

	}

	// Initial http client
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	tr := &http.Transport{TLSClientConfig: config}
	Client = &http.Client{Transport: tr}

}

func SplitName(name string) (city string, citycn string) {
	ns := strings.Split(name, "(")
	if len(ns) != 2 {
		log.Lx.WithFields(logrus.Fields{
			"name": name,
		}).Error("Input name wasn't matched with convention. eg: ", "Beijing (北京)")

		return name, ""
	}
	city = strings.Trim(ns[0], " ")
	citycn = strings.Trim(ns[1], ")")
	return city, citycn
}

func Copy2AirQuality(src OriginAirQuality) AirQuality {

	var dest AirQuality
	dest.StationIndex = src.Data.StationIndex
	dest.AQI = src.Data.AQI
	c, cn := SplitName(src.Data.City.Name)
	dest.City = c
	dest.CityCN = cn
	dest.Latitude = strconv.FormatFloat(src.Data.City.Geo[0], 'g', 6, 64)
	dest.Longitude = strconv.FormatFloat(src.Data.City.Geo[1], 'g', 6, 64)

	dest.Co = strconv.FormatFloat(src.Data.IAQI.Co.V, 'g', 6, 64)
	dest.H = strconv.FormatFloat(src.Data.IAQI.H.V, 'g', 6, 64)
	dest.No2 = strconv.FormatFloat(src.Data.IAQI.No2.V, 'g', 6, 64)
	dest.O3 = strconv.FormatFloat(src.Data.IAQI.O3.V, 'g', 6, 64)
	dest.P = strconv.FormatFloat(src.Data.IAQI.P.V, 'g', 6, 64)
	dest.Pm10 = strconv.FormatFloat(src.Data.IAQI.Pm10.V, 'g', 6, 64)
	dest.Pm25 = strconv.FormatFloat(src.Data.IAQI.Pm25.V, 'g', 6, 64)
	dest.So2 = strconv.FormatFloat(src.Data.IAQI.So2.V, 'g', 6, 64)
	dest.T = strconv.FormatFloat(src.Data.IAQI.T.V, 'g', 6, 64)
	dest.W = strconv.FormatFloat(src.Data.IAQI.W.V, 'g', 6, 64)

	dest.S = src.Data.OriginTime.S
	dest.TZ = src.Data.OriginTime.TZ
	dest.V = src.Data.OriginTime.V

	dest.IndexCity = "" + dest.City + "_" + strconv.Itoa(dest.StationIndex)

	h := sha1.New()
	h.Write([]byte(dest.IndexCity + "_" + strconv.Itoa(dest.V)))
	dest.IndexCityVHash = hex.EncodeToString(h.Sum(nil))
	return dest

}

func HttpGet(ctx context.Context, url string) ([]byte, error) {

	log.Lx.WithFields(logrus.Fields{
		"url": url,
	}).Info("request to")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Lx.WithFields(logrus.Fields{
			"url":   url,
			"error": err,
		}).Error("Http.NewRequest() was failed")
		return nil, err
	}
	resp, err := Client.Do(req)
	if err != nil {
		log.Lx.WithFields(logrus.Fields{
			"url":   url,
			"error": err,
		}).Error("Client.Do() was failed")
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Lx.WithFields(logrus.Fields{
			"url":   url,
			"error": err,
		}).Error("Failed to read buffer")
		return nil, err
	}

	log.Lx.WithFields(logrus.Fields{
		"url":  url,
		"byte": body,
	}).Debug("response from http server")
	return body, nil
}

func CityByIP(ctx context.Context, ip string) (string, error) {

	var ipStack IpStack
	url := IpStackServer + "/" + ip + "?access_key=" + IpStackServerToken
	buf, err := HttpGet(ctx, url)
	if err != nil {
		return "", err
	} else {
		err = json.Unmarshal(buf, &ipStack)
	}

	return ipStack.City, err

}

func convertAir(content []byte) (AirQuality, error) {
	var originAir OriginAirQuality
	var newAir AirQuality
	var apiError ApiError

	err := json.Unmarshal(content, &originAir)

	if originAir.Status == "error" {
		err = json.Unmarshal(content, &apiError)
		log.Lx.WithFields(logrus.Fields{
			"error": apiError,
		}).Error("return error from AQI service")
		return newAir, err

	}
	newAir = Copy2AirQuality(originAir)

	return newAir, nil

}

func AirbyCity(ctx context.Context, city string) (AirQuality, error) {

	url := AQIServer + "/feed/" + city + "/?token=" + AQIServerToken
	log.Lx.WithFields(logrus.Fields{
		"url": url,
	}).Info("request to AQI service")
	buf, err := HttpGet(ctx, url)
	if err != nil {
		return AirQuality{}, err
	}
	air, err := convertAir(buf)
	log.Lx.WithFields(logrus.Fields{
		"url": url,
		"air": air,
	}).Info("curated response from AQI service")
	return air, err
}

// Air Quality Index scale as defined by the US-EPA 2016 standard
func Readme() map[string]interface{} {
	var jsonMap map[string]interface{}
	str := `

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
					"Caution": "Active children and adults, and people with respiratory disease, such as asthma, should avoid prolonged outdoor exertion;everyone else, especially children, should limit prolonged outdoor exertion"
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
	`

	err := json.Unmarshal([]byte(str), &jsonMap)
	if err != nil {
		fmt.Println(err.Error())
	}
	return jsonMap
}
