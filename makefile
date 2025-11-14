build:
	@go build -o totion ./cmd/totion/main.go

run: build
	@./totion

clean:
	@rm -f totion

install:
	@go install ./cmd/totion

version:
	@echo "totion version 2.1.0"