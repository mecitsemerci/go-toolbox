.PHONY: lint format

lint:
	cd . && golangci-lint run --fix
	cd jsonconv && golangci-lint run --fix
	cd ratelimit && golangci-lint run --fix
	cd resilience && golangci-lint run --fix
	cd validation && golangci-lint run --fix

format:
	gofmt -s -w .