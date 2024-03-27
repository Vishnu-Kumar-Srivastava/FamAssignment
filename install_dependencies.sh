#!/bin/bash

# Check if Go is installed
if ! [ -x "$(command -v go)" ]; then
  echo "Error: Go is not installed. Please install Go before running this script."
  exit 1
fi

# Install Go dependencies
echo "Installing Go dependencies..."
go mod download

# Check if MongoDB is installed
if ! [ -x "$(command -v mongod)" ]; then
  echo "Warning: MongoDB is not installed. Some functionalities may not work without MongoDB."
fi

echo "Go dependencies installed successfully."
