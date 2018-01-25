PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: test

build:
	go build ./...

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

cover: test
	@echo ""
	@mkdir -p coverage/
	@echo "mode: set" > ./coverage/test.cov
	@for PKG in $(PACKAGES); do if [ -f $$GOPATH/src/$$PKG/coverage.out ]; then tail -q -n +2 $$GOPATH/src/$$PKG/coverage.out | grep -v '.pb.go' >> ./coverage/test.cov; fi; done;
	@go tool cover -func ./coverage/test.cov
	@go tool cover -html=./coverage/test.cov

