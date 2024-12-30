#!/bin/bash

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Build the application
echo "Building the TOPAYCHAIN application..."
go build -o TOPAYCHAIN ./cmd

if [ $? -eq 0 ]; then
    echo "Build successful. Application binary created: ./TOPAYCHAIN"
else
    echo "Build failed. Check errors and try again."
    exit 1
fi
