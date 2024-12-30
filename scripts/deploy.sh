#!/bin/bash

# Configuration
NODES=("127.0.0.1:8081" "127.0.0.1:8082" "127.0.0.1:8083")
APP_BINARY="./TOPAYCHAIN"

# Ensure the binary exists
if [ ! -f "$APP_BINARY" ]; then
    echo "Application binary not found. Run build.sh first."
    exit 1
fi

# Deploy nodes
for NODE in "${NODES[@]}"; do
    echo "Deploying node at $NODE..."
    # This assumes a deployment mechanism (e.g., scp, rsync)
    scp $APP_BINARY user@$NODE:/path/to/deploy/
    ssh user@$NODE "cd /path/to/deploy/ && ./TOPAYCHAIN &"
    echo "Node deployed at $NODE."
done

echo "Deployment completed."
