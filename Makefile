test:
	@CGO_ENABLED=0 go test -v ./...

build-image:
	docker build -t crew-lambda .
