.PHONY: all fmt-md

all:
	cd bookstore && make all
	cd bookstore2 && make all

fmt-md:
	@echo "Formatting root markdown files..."
	npx prettier --write '*.md'
	cd bookstore && make fmt-md
	cd bookstore2 && make fmt-md
