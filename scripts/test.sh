#!/bin/bash

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Run tests
echo "Running all tests..."
go test ./...

if [ $? -eq 0 ]; then
    echo "All tests passed successfully."
else
    echo "Some tests failed. Check the test logs above."
    exit 1
fi
