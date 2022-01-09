#!/bin/sh
curl -sSL https://git.io/g-install | sh -s bash -y
source /app/.bashrc
g install latest
echo "Running script"
go run scripts/augment.go
