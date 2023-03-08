fmt:
	@go fmt ./ ./api ./database ./internal

app:
	@go build -o build/bin/app ./

all: fmt app