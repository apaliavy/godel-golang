TOOLS += github.com/maxbrunsfeld/counterfeiter/v6

.PHONY: $(TOOLS)
$(TOOLS): %:
	GOBIN=$(GOBIN) go install $*

.PHONY: tools
tools: deps $(TOOLS)

mocks:
	@go generate ./...

test:
	go test --cover --race ./client/...
