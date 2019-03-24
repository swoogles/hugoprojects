package weather

import (
	"fmt"
	"os"
	"testing"
)

func TestGetBasicForecast(t *testing.T) {
	var darkSkyToken = os.Getenv("DARK_SKY_TOKEN")
	var forecast = GetBasicForecast(darkSkyToken)
	fmt.Println(forecast)
	t.Errorf("Failed forecast: %f", forecast.Currently.Temperature)
}
