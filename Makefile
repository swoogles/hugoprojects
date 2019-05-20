build:
	hugo -v
	mkdir -p lambdas
	# go get ./...
	go build -o lambdas/current_forecasts  ./functions/weather/current_forecasts