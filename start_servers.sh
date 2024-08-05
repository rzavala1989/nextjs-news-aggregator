#!/bin/bash

# Navigate to the Flask project directory
cd news_aggregator_recommendations

# Activate the virtual environment
source .venv/bin/activate

# Start the Flask server in the background
nohup python app.py &

# Wait for the Flask server to start
echo "Starting Flask server..."
sleep 3

# Navigate to the Go project directory
cd ../news-aggregator/cmd/server

# Start the Go server in the background
nohup go run main.go &

echo "Servers are running!"
