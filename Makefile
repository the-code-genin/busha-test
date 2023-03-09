fmt:
	@go fmt ./ ./api ./database ./internal ./swapi

app:
	@go build -o build/bin/app ./

all: fmt app