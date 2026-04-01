#!/bin/bash

# Prevent the script from continuing if a command fails
set -e

# Store colors for nice output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}▶ Starting Go Backend...${NC}"
cd user-service
# Run backend in the background
go run cmd/main.go &
BACKEND_PID=$!
cd ..

echo -e "${GREEN}▶ Starting Next.js Frontend...${NC}"
cd frontend/my-app
# Run frontend in the background
npm run dev &
FRONTEND_PID=$!
cd ../..

echo "========================================="
echo -e "${GREEN}Services are successfully running!${NC}"
echo "Backend (Go): PID $BACKEND_PID"
echo "Frontend (Next.js): PID $FRONTEND_PID"
echo "========================================="
echo "Press Ctrl+C to stop both."

# Clean exit: Kill both child processes when the script is stopped
trap "echo -e '\nStopping services...'; kill $BACKEND_PID $FRONTEND_PID 2>/dev/null; exit 0" SIGINT SIGTERM

# Keep the script alive and wait for processes
wait $BACKEND_PID $FRONTEND_PID
