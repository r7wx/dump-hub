all:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/dump-hub cmd/dump-hub/main.go
clean:
	rm -rf bin/