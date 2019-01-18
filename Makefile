.PHONY: build
build:
	go build -tags netgo -ldflags="-s -w" -o HaveIBeenKeePassed

.PHONY: mac
mac:
	GOOS="darwin" GOARCH="amd64" go build -tags netgo -ldflags="-s -w" -o HaveIBeenKeePassed

.PHONY: linux
linux:
	GOOS="linux" GOARCH="amd64" go build -tags netgo -ldflags="-s -w" -o HaveIBeenKeePassed

.PHONY: windows
windows:
	GOOS="windows" GOARCH="amd64" go build -tags netgo -ldflags="-s -w" -o HaveIBeenKeePassed.exe

.PHONY: clean
clean:
	rm -f HaveIBeenKeePassed HaveIBeenKeePassed.exe
