.PHONY : dep run

dep:
	@echo "installing dependence..."
	@go get -v && go mod tidy

run:
	@echo "running..."
	@go run main.go
