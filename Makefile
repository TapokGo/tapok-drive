.DEFAULT_GOAL := help
PROJECT_NAME := tapokdrive
BIN_DIR := bin
LOG_PATH := app.log
DEV_BRANCH := dev

.PHONY: test lint run sync help 

test:
	@echo --------------------
	@echo "Starting tests..."
	@go test -v ./...
	@echo "Test complete"

lint:
	@echo --------------------
	@echo "Starting linters"
	@golangci-lint run

run: lint test
	@echo --------------------
	@echo "Starting project..."
	@docker compose up --build

sync:
	@git stash -m "WIP changes before 'make sync'"
	@git fetch origin dev
	@git merge origin/dev
	@echo "Synced with origin/dev"
	@echo "Uncommited changes in stash (run 'git stash pop' to restore)"

help:
	@echo "Available commands:"
	@echo ""
	@echo "  test     — Run all Go tests with verbose output"
	@echo "  lint     — Run golangci-lint for code quality checks"
	@echo "  build    — Run tests and linter, then build binary to ./bin/tapokdrive"
	@echo "  run      — Build (if needed) and start the application"
	@echo "  clean    — Remove ./bin directory and app.log"
	@echo "  sync     — Stash uncommitted changes, fetch origin/dev, and merge into current branch"
	@echo ""
	@echo "Note: after 'make sync', restore your changes with 'git stash pop'." 
