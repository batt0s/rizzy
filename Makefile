# Build for linux and windows 64-bit
build:
	cd src; \
	env GOOS=linux GOARCH=amd64 go build -o ../bin/rizzy-linux-amd64 .; \
	env GOOS=windows GOARCH=amd64 go build -o ../bin/rizzy-windows-amd64.exe .;

# Test all
test:
	cd src; \
	go test ./lexer; \
	go test ./parser; \
	go test ./ast; \
	go test ./object; \
	go test ./evaluator;

# Run Temp
dev:
	cd src; go run .;

# Run binary
run:
	./bin/rizzy-linux-amd64