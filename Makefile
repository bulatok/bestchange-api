run:
	go build -o don && ./don
test:
	go test ./... -v