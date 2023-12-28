build:
	go build -o bin/app

run: tailwindcss templ build
	./bin/app

test:
	go test -v ./... -count=1 


tailwindcss:
	bun run tailwindcss --config tailwind.config.js -i input.css -o static/css/styles.css

templ:
	templ generate

