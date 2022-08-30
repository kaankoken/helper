build:
	go build main.go

lint:
	golint ./...

run-vendor:
	bash ./scripts/run_vendor.sh

run-vet:
	go vet ./...

static-check:
	staticcheck ./...

test:
	bash ./scripts/run_tests.sh
