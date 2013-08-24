all:
	@mkdir -p bin/
	@go build -o bin/httphead httphead.go
