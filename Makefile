go-compile: go-clean go-get go-build

go-build:
	@echo "  >  Building binary..."
	@go build -o vcloud-inet

go-run:
	@echo "  >  Running application..."
	@go run main.go parser.go render.go
