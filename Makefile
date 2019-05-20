build:
	mkdir -p lambdas
	go get ./...
	go build -o lambdas/current_forecasts  ./functions/weather
	hugo -v
