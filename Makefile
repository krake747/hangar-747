.PHONY: all fmt-md

all:
	make fmt-md

fmt-md:
	@echo "Formatting all markdown files..."
	npx prettier --write '**/*.md'