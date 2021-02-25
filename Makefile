.PHONY: \
	all build build-release build-debug install uninstall

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

--install:
	@echo "Install raspichamber service"
	cp -f .bin/raspichamber /usr/sbin/
	cp -f system/systemd/raspichamber.service /etc/systemd/system/
	systemctl enable raspichamber
	systemctl start raspichamber

	@echo "Install raspichamber_display service"
	cp -f python/raspichamber_display.py /usr/sbin/
	cp -f system/systemd/raspichamber_display.service /etc/systemd/system/
	systemctl enable raspichamber_display
	systemctl start raspichamber_display

--uninstall:
	@echo "Uninstall raspichamber service"
	systemctl stop raspichamber
	systemctl disable raspichamber
	rm -f /etc/systemd/system/raspichamber.service
	rm -f /usr/sbin/raspichamber

	@echo "Uninstall raspichamber_display service"
	systemctl stop raspichamber_display
	systemctl disable raspichamber_display
	rm -f /etc/systemd/system/raspichamber_display.service
	rm -f /usr/sbin/raspichamber_display.py