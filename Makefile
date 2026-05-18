.DEFAULT_GOAL = check

include lab.mk
-include .github/ci.env

# Run all checks required to validate the codebase before merging.
.PHONY: check
check: test lint coverage-check

# Install Git hooks for automatic testing and linting.
.PHONY: githooks
githooks:
	@git config core.hooksPath .githooks
	@chmod +x .githooks/pre-commit .githooks/pre-push
	@echo "✅ Git hooks installed successfully"

# Run all tests in the project.
.PHONY: test
test:
	@$(MAKE) unit-test || { \
		echo "❌ Unit tests failed; skipping integration tests"; \
		exit 1; \
	}
	@$(MAKE) integration-test

# Run unit tests.
.PHONY: unit-test
unit-test:
	@echo 🧪 Running unit tests...
	@mkdir -p tmp
	@cd autograder && ( \
		pkgs=$$(go list ./... | grep -v 'internal/test/integration' | grep -v 'internal/test/fake'); \
		if [ -n "$$pkgs" ]; then \
			go test -race -count=1 \
			-covermode=atomic \
			-coverprofile=../tmp/coverage_unit.out \
			$$pkgs; \
		else \
			echo "no packages to test"; \
		fi \
	)

# Run integration tests.
.PHONY: integration-test
integration-test:
	@mkdir -p tmp
	@echo "🐳 Starting containers..."
	@set -e; \
	gomodcache=$$(cd autograder && go env GOMODCACHE); \
	gocache=$$(cd autograder && go env GOCACHE); \
	trap 'echo "🐳 Stopping containers..."; GO_TEST_GOMODCACHE_HOST="'"$$gomodcache"'" GO_TEST_GOCACHE_HOST="'"$$gocache"'" docker compose -f docker-compose.test.yml --env-file .env.test down' EXIT; \
	GO_TEST_GOMODCACHE_HOST="$$gomodcache" GO_TEST_GOCACHE_HOST="$$gocache" docker compose -f docker-compose.test.yml --env-file .env.test up -d redis-test mongo-test cassandra-test neo4j-test; \
	echo 🧪 Running integration tests...; \
	GO_TEST_GOMODCACHE_HOST="$$gomodcache" GO_TEST_GOCACHE_HOST="$$gocache" docker compose -f docker-compose.test.yml --env-file .env.test run --rm test-runner sh -c '\
		coverpkgs=$$(go list ./... | grep -v "internal/test/integration" | grep -v "internal/test/fake" | tr "\\n" ","); \
		go test -tags=integration \
		-race -count=1 -v \
		-covermode=atomic \
		-coverprofile=../tmp/coverage_integration.out \
		-coverpkg=$$coverpkgs \
		./internal/test/integration/...'

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


# Generate coverage report combining unit and integration tests.
# Outputs total coverage percentage to tmp/coverage_total.out,
# covered lines to tmp/coverage_*.out
# and uncovered lines to tmp/uncovered.out.
.PHONY: coverage
coverage:
	@echo 🧪 Calculating test coverage...
	@sh scripts/coverage/report.sh

# Check if the total test coverage meets the defined threshold.
.PHONY: coverage-check
coverage-check:
	@echo 🧪 Checking test coverage...
	@COVERAGE_THRESHOLD=$(COVERAGE_THRESHOLD) sh scripts/coverage/check.sh
