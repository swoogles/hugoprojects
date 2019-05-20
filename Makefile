build:
	hugo -v
	mkdir -p functions
	# go get ./...
	go build -o functions/hello-lambda  ./lambdas/weather/current_forecasts/