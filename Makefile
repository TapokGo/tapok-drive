.DEFAULT_GOAL := help
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
	@git fetch origin ${DEV_BRANCH}
	@git merge origin/${DEV_BRANCH}
	@echo "Synced with origin/${DEV_BRANCH}"
	@echo "Uncommited changes in stash (run 'git stash pop' to restore)"

help:
	@echo "Available commands:"
	@echo ""
	@echo "  test     — Run all Go tests with verbose output"
	@echo "  lint     — Run golangci-lint for code quality checks"
	@echo "  run      — Start the application"
	@echo "  sync     — Stash uncommitted changes, fetch origin/${DEV_BRANCH}, and merge into current branch"
	@echo ""
	@echo "Note: after 'make sync', restore your changes with 'git stash pop'." 
