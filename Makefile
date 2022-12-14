OS = linux

default: dev

dev: start-db
	find . -name '*.go' | entr -cr go run .

start-server: start-db
	GIN_MODE=release bin/openmusic-go-${OS}

clean: 
	rm -f bin/*

build: clean
	GOOS=${OS} GOARCH=amd64 go build -o bin/openmusic-go-${OS} main.go

cross-compile: clean
	GOOS=linux GOARCH=amd64 go build -o bin/openmusic-go-linux main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/openmusic-go-darwin main.go
	GOOS=windows GOARCH=amd64 go build -o bin/openmusic-go-windows.exe main.go

start-db:
	systemctl is-active postgresql || sudo systemctl start postgresql

stop-db:
	sudo systemctl stop postgresql

status-db:
	systemctl status postgresql
