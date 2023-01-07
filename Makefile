.PHONY: run
run:
	go run cmd/av-random/main.go

.PHONY: vet
vet:
	go vet ./...