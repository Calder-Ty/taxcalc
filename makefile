help:
	@echo "Available commands"
	@echo "	clean:		Clean from last build"
	@echo "	build:		Build new version"


clean:
	rm -r ./bin/*

build: clean
	go build -o bin/taxcalc src/main.go
