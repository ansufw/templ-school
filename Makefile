.PHONY: build dev clean templ run

# Build the application
build: templ
	go build -o bin/app cmd/web/main.go

# Development mode with hot reload
dev:
	@echo "Starting development server..."
	@air -c .air.toml || go run cmd/web/main.go

# Generate templ files
templ:
	@echo "Generating templ files..."
	@templ generate --watch --proxy="http://localhost:8080"

# download tailwind
tw:
	@echo "download tailwind"
	@chmod +x ./scripts/download-tailwind.sh
	./scripts/download-tailwind.sh


# Run the application
run: build
	./bin/app

# Clean build artifacts
clean:
	rm -rf bin/
	find . -name "*_templ.go" -delete