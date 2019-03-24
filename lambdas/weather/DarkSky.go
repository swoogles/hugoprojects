package  weather

import (
	"encoding/json"
	"net/http"
	"time"
)

type ForeCast struct {
	Timezone string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetBasicForecast(darkSkyToken string) ForeCast {
	req3, _ := http.NewRequest("GET", "https://api.darksky.net/forecast/" + darkSkyToken + "/37.8267,-122.4233", nil)
	resp3, _ := myClient.Do(req3)
	defer resp3.Body.Close()
	var weatherForecast ForeCast
	json.NewDecoder(resp3.Body).Decode(&weatherForecast)
	return weatherForecast;
}

