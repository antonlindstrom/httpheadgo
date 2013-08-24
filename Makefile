all:
	@mkdir -p bin/
	@go build -o bin/httphead httphead.go

clean:
	@rm -rf bin/
