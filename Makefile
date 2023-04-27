build:
	go build -o target/storageid cmd/main.go

clean:
	rm -rf target/

test:
	go test ./...