.PHONY: \
	all build build-release build-debug

--build:
	@echo "Build raspichamber"
	@echo $(BUILD_ARGS)
	go build $(BUILD_ARGS) -o ".bin/raspichamber" cmd/raspichamber/main.go

build: build-debug

build-release: BUILD_TYPE = release
build-release: --build

build-debug: BUILD_TYPE = debug
build-debug: BUILD_ARGS += -tags pprof
build-debug: --build

build-profile: BUILD_TYPE = profile
build-profile: BUILD_ARGS += -tags pprof
build-profile: --build
