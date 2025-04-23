all: sops-sample

.PHONY: sops-sample   # leave dependency handling to go
sops-sample:
	go build ./...
