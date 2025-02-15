# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: geth android ios evm all test clean rocksdb
.PHONY: gwemix-linux

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

# USE_ROCKSDB
# - undefined | "NO": Do not use
# - "YES": build a static lib from rocksdb directory, and use that one
# - "EXISTING": use existing rocksdb shared lib.
ifndef USE_ROCKSDB
  ifeq ($(shell uname), Linux)
    USE_ROCKSDB = YES
  else
    USE_ROCKSDB = NO
  endif
endif
ifneq ($(shell uname), Linux)
  USE_ROCKSDB = NO
endif

ifneq ($(USE_ROCKSDB), NO)
ROCKSDB_DIR=$(shell pwd)/rocksdb
ROCKSDB_TAG=-tags rocksdb
endif

gwemix.tar.gz: gwemix logrot
	@[ -d build/conf ] || mkdir -p build/conf
	@cp -p wemix/scripts/gwemix.sh build/bin/
	@cp -p wemix/scripts/config.json.example		\
		wemix/scripts/genesis-template.json		\
		build/conf/
	@(cd build; tar cfz gwemix.tar.gz bin conf)
	@echo "Done building build/gwemix.tar.gz"

gwemix: rocksdb 
ifeq ($(USE_ROCKSDB), NO)
	$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/gwemix
else
	CGO_CFLAGS=-I$(ROCKSDB_DIR)/include \
		CGO_LDFLAGS="-L$(ROCKSDB_DIR) -lrocksdb -lm -lstdc++ $(shell awk '/PLATFORM_LDFLAGS/ {sub("PLATFORM_LDFLAGS=", ""); print} /JEMALLOC=1/ {print "-ljemalloc"}' < $(ROCKSDB_DIR)/make_config.mk)" \
		$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/gwemix
endif
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gwemix\" to launch gwemix."

logrot:
	$(GORUN) build/ci.go install ./cmd/logrot

geth:
	$(GORUN) build/ci.go install ./cmd/geth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geth\" to launch geth."

dbbench: rocksdb
ifeq ($(USE_ROCKSDB), NO)
	$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/dbbench
else
	CGO_CFLAGS=-I$(ROCKSDB_DIR)/include \
		CGO_LDFLAGS="-L$(ROCKSDB_DIR) -lrocksdb -lm -lstdc++ $(shell awk '/PLATFORM_LDFLAGS/ {sub("PLATFORM_LDFLAGS=", ""); print} /JEMALLOC=1/ {print "-ljemalloc"}' < $(ROCKSDB_DIR)/make_config.mk)" \
		$(GORUN) build/ci.go install $(ROCKSDB_TAG) ./cmd/dbbench
endif

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/geth.aar\" to use the library."
	@echo "Import \"$(GOBIN)/geth-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

test-short: all
	$(GORUN) build/ci.go test -short

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/* build/conf
	@ROCKSDB_DIR=$(ROCKSDB_DIR);			\
	if [ -e $${ROCKSDB_DIR}/Makefile ]; then	\
		cd $${ROCKSDB_DIR};			\
		make clean;				\
	fi

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

gwemix-linux:
	@docker --version > /dev/null 2>&1;				\
	if [ ! $$? = 0 ]; then						\
		echo "Docker not found.";				\
	else								\
		docker build -t wemix/builder:local			\
			-f Dockerfile.wemix . &&			\
		docker run -e HOME=/tmp --rm -v $(shell pwd):/data	\
			-w /data wemix/builder:local			\
			"git config --global --add safe.directory /data;\
			 make USE_ROCKSDB=$(USE_ROCKSDB)";		\
	fi

ifneq ($(USE_ROCKSDB), YES)
rocksdb:
else
rocksdb:
	@[ ! -e rocksdb/.git ] && git submodule update --init rocksdb;	\
	cd $(ROCKSDB_DIR) && PORTABLE=1 make -j8 static_lib;
endif
