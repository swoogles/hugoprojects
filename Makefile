build:
	mkdir -p lambdas
	go get ./functions/...
	go build -o lambdas/current_forecasts  ./functions/weather/current_forecasts
	hugo -v
