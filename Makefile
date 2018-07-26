APP=srcon
VERSION=1.2.0

go get -u ./...

all: dist/$(APP)-$(VERSION)-linux.zip dist/$(APP)-$(VERSION)-darwin.zip dist/$(APP)-$(VERSION)-windows.zip

dist/$(APP)-$(VERSION)-linux.zip: build/linux/$(APP) dist
	cd build/linux; zip ../../dist/$(APP)-$(VERSION)-linux.zip $(APP)

dist/$(APP)-$(VERSION)-darwin.zip: build/darwin/$(APP) dist
	cd build/darwin; zip ../../dist/$(APP)-$(VERSION)-darwin.zip $(APP)

dist/$(APP)-$(VERSION)-windows.zip: build/windows/$(APP).exe dist
	cd build/windows; zip ../../dist/$(APP)-$(VERSION)-windows.zip $(APP).exe

dist:
	mkdir -p dist

build/linux/$(APP): cmd/$(APP)/*.go
	GOOS=linux GOARCH=amd64 go build -o build/linux/$(APP) cmd/$(APP)/*.go

build/darwin/$(APP): cmd/$(APP)/*.go
	GOOS=darwin GOARCH=amd64 go build -o build/darwin/$(APP) cmd/$(APP)/*.go

build/windows/$(APP).exe: cmd/$(APP)/*.go
	GOOS=windows GOARCH=amd64 go build -o build/windows/$(APP).exe cmd/$(APP)/*.go

clean:
	rm -rf build/* dist/*
