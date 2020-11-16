ifeq ($(CI),)
	# Not in CI
	VERSION := dev-$(USER)
	VERSION_HASH := $(shell git rev-parse HEAD | cut -b1-8)
else
	# In CI
	ifneq ($(GITHUB_REF),)
		VERSION := $(GITHUB_REF)
	else
		# No tag, so make one
		VERSION := $(shell git describe --tags 2>/dev/null)
	endif
	VERSION_HASH := $(shell echo "$(GITHUB_SHA)" | cut -b1-8)
endif

install:
	go install -v -ldflags "-X code.vegaprotocol.io/go-wallet/cmd.Version=${VERSION} -X code.vegaprotocol.io/go-wallet/cmd.VersionHash=${VERSION_HASH}"

proto:
	protoc --go_out=paths=source_relative,plugins=grpc:. ./proto/*.proto
	protoc --go_out=paths=source_relative,plugins=grpc:. ./proto/api/*.proto


release-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o build/vegawallet-windows-amd64 -ldflags "-X code.vegaprotocol.io/go-wallet/cmd.Version=${VERSION} -X code.vegaprotocol.io/go-wallet/cmd.VersionHash=${VERSION_HASH}"

release-macos:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -o build/vegawallet-darwin-amd64 -ldflags "-X code.vegaprotocol.io/go-wallet/cmd.Version=${VERSION} -X code.vegaprotocol.io/go-wallet/cmd.VersionHash=${VERSION_HASH}"

release-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o build/vegawallet-linux-amd64 -ldflags "-X code.vegaprotocol.io/go-wallet/cmd.Version=${VERSION} -X code.vegaprotocol.io/go-wallet/cmd.VersionHash=${VERSION_HASH}"


.PHONY: proto
