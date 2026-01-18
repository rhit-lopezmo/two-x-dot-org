DEFAULT_GOAL: run-with-air

build: tidy gen
	go build -o build/my-website main.go

run:
	@./build/my-website

run-with-air:
	air -c air.toml

gen-and-watch:
	go tool templ generate --watch

render-build:
	go tool templ generate
	mkdir -p build
	go build -tags netgo -ldflags '-s -w' -o build/app

clean:
	rm -rf build
	mkdir build

tidy:
	go mod tidy
