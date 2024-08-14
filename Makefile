SERVICE_NAME = websocket
ENTRY_FILE = websocket.go
HOOKS_DIR = .git/hooks
PRE_COMMIT_HOOK = $(HOOKS_DIR)/pre-commit

.PHONY: all
all: build

.PHONY: build
build:
	@echo "Building the application..."
	GO111MODULE=on go build -o bin/$(SERVICE_NAME) $(CMD_DIR)

.PHONY: test
test:
	@echo "Running tests..."
	GO111MODULE=on go test ./...

.PHONY: clean
clean:
	@echo "Cleaning the build..."
	-rm -f bin/$(SERVICE_NAME)

.PHONY: fmt
fmt:
	@echo "Formatting the code..."
	GO111MODULE=on go fmt ./...

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	GO111MODULE=on go mod tidy

.PHONY: prepare
prepare: $(PRE_COMMIT_HOOK)

$(PRE_COMMIT_HOOK):
	@echo "Creating pre-commit hook..."
	@mkdir -p $(HOOKS_DIR)
	@echo '#!/bin/sh' > $(PRE_COMMIT_HOOK)
	@echo '' >> $(PRE_COMMIT_HOOK)
	@echo '# Run tests' >> $(PRE_COMMIT_HOOK)
	@echo 'echo "Running tests..."' >> $(PRE_COMMIT_HOOK)
	@echo 'make test' >> $(PRE_COMMIT_HOOK)
	@echo 'TEST_RESULT=$$?' >> $(PRE_COMMIT_HOOK)
	@echo '' >> $(PRE_COMMIT_HOOK)
	@echo 'if [ $$TEST_RESULT -ne 0 ]; then' >> $(PRE_COMMIT_HOOK)
	@echo '    echo "Tests failed. Aborting commit."' >> $(PRE_COMMIT_HOOK)
	@echo '    exit 1' >> $(PRE_COMMIT_HOOK)
	@echo 'fi' >> $(PRE_COMMIT_HOOK)
	@echo '' >> $(PRE_COMMIT_HOOK)
	@echo 'fi' >> $(PRE_COMMIT_HOOK)
	@echo '' >> $(PRE_COMMIT_HOOK)
	@echo 'echo "All checks passed. Proceeding with commit."' >> $(PRE_COMMIT_HOOK)
	@echo 'exit 0' >> $(PRE_COMMIT_HOOK)
	@chmod +x $(PRE_COMMIT_HOOK)
	@echo "Pre-commit hook created successfully."

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make          Build the application"
	@echo "  make test     Run tests"
	@echo "  make clean    Clean the build"
	@echo "  make fmt      Format the code"
	@echo "  make deps     Install dependencies"
	@echo "  make help     Display this help message"