package collector

import (
	"context"
	"github.com/sn-2/weather-forecast/model"
)

type WeatherService interface {
	Forecast(ctx context.Context, city string, days int) ([]model.Weather, error)
}
