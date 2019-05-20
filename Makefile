build:
	hugo -v
	mkdir -p functions
	go get ./...
	go build -o functions/current_forecasts  ./lambdas/weather/current_forecasts