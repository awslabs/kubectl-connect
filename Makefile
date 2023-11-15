export GO111MODULE=on

.PHONY: test
test:
	go test ./cmd/... -coverprofile cover.out

.PHONY: bin
bin: fmt vet
	go build -o bin/kubectl-connect github.com/awslabs/kubectl-connect/cmd/plugin

.PHONY: fmt
fmt:
	go fmt ./cmd/...

.PHONY: vet
vet:
	go vet ./cmd/...

.PHONY: install
install: bin
	cp bin/kubectl-connect /usr/local/bin/kubectl-connect

.PHONY: clean
clean:
	rm -rf bin

