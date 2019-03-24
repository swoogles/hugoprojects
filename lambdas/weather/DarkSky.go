package  weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type DataPoint struct {
	ApparentTemperature float64
	Temperature float64
	Summary string
	CloudCover float64
	PrecipIntensity float64
	PrecipProbability float64
}

type ForeCast struct {
	Timezone string
	Currently DataPoint
}

var myClient = &http.Client{Timeout: 10 * time.Second}

//func getForeCastJson(url string) ForeCast {
//}

func GetBasicForecast(darkSkyToken string) ForeCast {
	req3, _ := http.NewRequest("GET", "https://api.darksky.net/forecast/" + darkSkyToken + "/38.8697,-106.9878", nil)
	resp3, error := myClient.Do(req3)
	if (error != nil) {
		fmt.Fprintf(os.Stderr, "error: %v\n", error)
		os.Exit(1)
	}

	defer resp3.Body.Close()
	var weatherForecast ForeCast
	json.NewDecoder(resp3.Body).Decode(&weatherForecast)
	return weatherForecast;
}

