VERSION=0.1.0

.PHONY: clean deps build

clean:
	rm .bands.yaml
	rm bands

deps:
	go get ./...

build: deps
	GOOS=darwin go build -o ./build/bands-osx-${VERSION} main.go
	GOOS=windows go build -o ./build/bands-win-${VERSION} main.go
	GOOS=linux go build -o ./build/bands-linux-${VERSION} main.go
	aws s3 sync ./build/ s3://bands-sh/releases
	aws s3 cp ./scripts/get-bands.sh s3://bands-sh/get-bands.sh