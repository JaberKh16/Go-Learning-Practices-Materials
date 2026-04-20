#!/bin/bash

# default config path (optional)
# DEFAULT_CONFIG="./config/config.yaml"

# take argument or fallback to default
# CONFIG_PATH=${1:-$DEFAULT_CONFIG}

# echo "Starting Go server with config: $CONFIG_PATH"

# go run app.go -config="$CONFIG_PATH"
#
#



#!/bin/bash

set -e  # stop on error

APP_NAME="go-server"
CONFIG_PATH=${1:-"./config.yaml"}

if [ ! -f "$CONFIG_PATH" ]; then
  echo "❌ Config file not found: $CONFIG_PATH"
  exit 1
fi

echo "✅ Config found: $CONFIG_PATH"

go build -o $APP_NAME main.go

echo "🚀 Starting server..."
./$APP_NAME -config="$CONFIG_PATH"
