OS = linux

default: dev

dev: start-db
	reflex -r "\.go$$" -s -- go run -tags=jsoniter main.go

start-server: build start-db
	GIN_MODE=release bin/openmusic-go-${OS}

clean: 
	rm -f bin/*

build: clean
	GOOS=${OS} GOARCH=amd64 go build -o bin/openmusic-go-${OS} -tags=jsoniter main.go

cross-compile: clean
	GOOS=linux GOARCH=amd64 go build -o bin/openmusic-go-linux -tags=jsoniter main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/openmusic-go-darwin -tags=jsoniter main.go
	GOOS=windows GOARCH=amd64 go build -o bin/openmusic-go-windows.exe -tags=jsoniter main.go

start-db:
	systemctl is-active postgresql || sudo systemctl start postgresql

stop-db:
	sudo systemctl stop postgresql

status-db:
	systemctl status postgresql