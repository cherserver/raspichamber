.PHONY: \
	all build build-release build-debug

--build:
	@echo "Build binaries"
	@echo $(BUILD_ARGS)
	@ERR=0; \
	for CMD in $$(find "./cmd" -maxdepth 1 -mindepth 1 -type d -print); do \
		echo "Build $$(basename "$${CMD}")"; \
		BIN=$$(basename "$${CMD}"); \
		go build $(BUILD_ARGS) -o "$(GO_DIR)/.bin/$${BIN}/$(BUILD_TYPE)/$${BIN}" "$${CMD}" || { \
			ERR=$$?; \
			break; \
		}; \
	done; \
	if [ $$ERR != 0 ]; then \
		exit $$ERR; \
	fi

build: build-debug

build-release: BUILD_TYPE = release
build-release: --build

build-debug: BUILD_TYPE = debug
build-debug: BUILD_ARGS += -race -tags pprof
build-debug: --build

build-profile: BUILD_TYPE = profile
build-profile: BUILD_ARGS += -tags pprof
build-profile: --build
