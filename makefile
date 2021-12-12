WHAT := autoOrder

PWD ?= $(shell pwd)

VERSION   ?= $(shell git describe --tags --always)
REVISION  ?= $(shell git rev-parse HEAD)
BRANCH    ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDUSER ?= $(shell id -un)
BUILDTIME ?= $(shell date '+%Y%m%d-%H:%M:%S')

.PHONY: build build-darwin-amd64 build-linux-amd64 build-windows-amd64 clean release static

build: static
	for target in $(WHAT); do \
		go build -ldflags "-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Version=${VERSION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Revision=${REVISION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Branch=${BRANCH} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/$$target ./cmd/$$target; \
	done

build-darwin-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -a -installsuffix cgo -ldflags "-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Version=${VERSION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Revision=${REVISION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Branch=${BRANCH} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/autoOrder-${VERSION}-darwin-amd64/$$target ./cmd/$$target; \
	done

build-linux-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -installsuffix cgo -ldflags "-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Version=${VERSION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Revision=${REVISION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Branch=${BRANCH} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/autoOrder-${VERSION}-linux-amd64/$$target ./cmd/$$target; \
	done

build-windows-amd64:
	for target in $(WHAT); do \
		CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -a -installsuffix cgo -ldflags "-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Version=${VERSION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Revision=${REVISION} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.Branch=${BRANCH} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildUser=${BUILDUSER} \
			-X github.com/furkansahinfs/AutoOrder-Backend/pkg/version.BuildDate=${BUILDTIME}" \
			-o ./bin/autoOrder-${VERSION}-windows-amd64/$$target.exe ./cmd/$$target; \
	done

clean:
	rm -rf ./bin;

prepare-release: clean vue-build static
	$(info "Run: 'git commit' and 'git tag', then 'make release'")

release: build-darwin-amd64 build-linux-amd64 build-windows-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/autoOrder-${VERSION}-darwin-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/autoOrder-${VERSION}-linux-amd64
	cp ${PWD}/LICENSE ${PWD}/bin/autoOrder-${VERSION}-windows-amd64
	cp ${PWD}/config.yml ${PWD}/bin/autoOrder-${VERSION}-darwin-amd64
	cp ${PWD}/config.yml ${PWD}/bin/autoOrder-${VERSION}-linux-amd64
	cp ${PWD}/config.yml ${PWD}/bin/autoOrder-${VERSION}-windows-amd64
	cd ${PWD}/bin; tar cfvz autoOrder-${VERSION}-darwin-amd64.tar.gz ./autoOrder-${VERSION}-darwin-amd64
	cd ${PWD}/bin; tar cfvz autoOrder-${VERSION}-linux-amd64.tar.gz ./autoOrder-${VERSION}-linux-amd64
	cd ${PWD}/bin; tar cfvz autoOrder-${VERSION}-windows-amd64.tar.gz ./autoOrder-${VERSION}-windows-amd64

static:
	go build -o esc esc.go
	./esc -o pkg/static/static.go -pkg static