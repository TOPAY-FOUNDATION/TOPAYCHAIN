#!/bin/bash

# Check if the application binary exists
if [ ! -f "./TOPAYCHAIN" ]; then
    echo "Application binary not found. Run build.sh first."
    exit 1
fi

# Run the application
echo "Starting the TOPAYCHAIN application..."
./TOPAYCHAIN
