GOPATH="$(shell echo $$GOPATH)"
BINDIR = bin

BINFILE=go

APPS = demo_app
all: $(APPS)

$(APPS): $(BINDIR)
	@echo "build $@"
	@export
	@echo "gopath $(GOPATH)"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on GOPATH=$(GOPATH) $(BINFILE) build -o $(BINDIR)/$@ main.go

test:
	@GOPATH=$(GOPATH) go test ./...

$(BINDIR):
	mkdir -p $(BINDIR)
