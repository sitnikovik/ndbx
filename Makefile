.DEFAULT_GOAL = check

# Run all checks required to validate the codebase before merging.
.PHONY: check
check: test lint

# Install Git hooks for automatic testing and linting.
.PHONY: githooks
githooks:
	@git config core.hooksPath .githooks
	@chmod +x .githooks/pre-commit .githooks/pre-push
	@echo "✅ Git hooks installed successfully"

# Run all tests in the project.
.PHONY: test
test:
	@cd autograder && go test ./... -race -count=1

# Lint the codebase.
.PHONY: lint
lint:
	@cd autograder && golangci-lint run ./...

# Remove "-dev" suffix from all files.
.PHONY: .rmdev
.rmdev:
	@if git remote get-url origin | grep -q '\-dev\.git$$'; then \
		echo "❌ Error: Cannot run .rmdev in a repository with '-dev' in its name"; \
		echo "Current remote: $$(git remote get-url origin)"; \
		exit 1; \
	fi
	@echo "Removing '-dev' from all files..."
	@grep -rl -E -- '-dev' . \
		--exclude-dir=.git \
		--exclude-dir=node_modules \
		--exclude-dir=vendor 2>/dev/null | grep -v 'Makefile' | while read file; do \
		sed -e 's/-dev//g' "$$file" > "$$file.tmp" && mv "$$file.tmp" "$$file"; \
	done
	@echo "✅ Done"