.PHONY: run
run:
	go run cmd/av-random/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./... -v -count=1 | grep -iv 'no test files'