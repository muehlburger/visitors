build:
	go build -o bin/visitors ./cmd/visitors

run:
	go run ./cmd/visitors

compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/visitors-freebsd-386 ./cmd/visitors
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/visitors-linux-386 ./cmd/visitors
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/visitors-386.exe ./cmd/visitors
    # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/visitors-freebsd-amd64 ./cmd/visitors
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/visitors-darwin-amd64 ./cmd/visitors
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/visitors-linux-amd64 ./cmd/visitors
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/visitors-amd64.exe ./cmd/visitors
