# Define the binary name
BINARY=./build/awktutor

# Define the Go files
GOFILES=$(wildcard *.go)

# Default target to build the project
all: build

# Build the project
build:
	go build -o $(BINARY) $(GOFILES)

# Run the project
run: build
	./$(BINARY)

# Clean up binary
clean:
	rm -f $(BINARY)

# Rebuild the project
rebuild: clean build
