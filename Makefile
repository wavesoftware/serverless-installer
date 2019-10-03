PROJECT_DIR = "./"
CLI_DIR     = "$(PROJECT_DIR)/cmd/installer"
BIN         = "$(PROJECT_DIR)/bin/serverless-installer"

GO           ?= go
GOLINT       ?= $(GO)lint
RICHGO       ?= rich$(GO)

RESET          = \033[0m
make_std_color = \033[3$1m      # defined for 1 through 7
make_color     = \033[38;5;$1m  # defined for 1 through 255
BLUE = $(strip $(call make_color,44))
PINK = $(strip $(call make_color,210))
RED = $(strip $(call make_color,206))
GREEN = $(strip $(call make_color,120))
DGREEN = $(strip $(call make_color,106))
GRAY = $(strip $(call make_color,224))

.PHONY: default
default: binary

.PHONY: builddeps
builddeps:
	@$(GO) get github.com/kyoh86/richgo
	@$(GO) get golang.org/x/lint/golint

.PHONY: clean
clean: builddeps
	@echo " $(GRAY)🛁 Cleaning$(RESET)"
	@rm -fv $(BIN)

.PHONY: check
check: builddeps
	@echo " $(PINK)🛂 Checking$(RESET)"
	$(GOLINT) -set_exit_status ./...

.PHONY: test
test: check
	@echo " $(GREEN)✔️ Testing$(RESET)"
	$(RICHGO) test -cover ./...

.PHONY: binary
binary: test
	@echo " $(BLUE)🔨 Building$(RESET)"
	$(RICHGO) build -o $(BIN) $(CLI_DIR)

.PHONY: run
run: binary
	@echo " $(RED)🏃 Running$(RESET)"
	$(BIN) $(args)
