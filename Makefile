PROJECT_DIR = $(shell readlink -f .)
CLI_DIR     = "$(PROJECT_DIR)/cmd/installer"
BUILD_DIR   = "$(PROJECT_DIR)/build"
BIN         = "$(BUILD_DIR)/serverless-installer"

GO           ?= go
GOLINT       ?= $(GO)lint
RICHGO       ?= rich$(GO)

RESET          = \033[0m
make_std_color = \033[3$1m      # defined for 1 through 7
make_color     = \033[38;5;$1m  # defined for 1 through 255
BLUE = $(strip $(call make_color,38))
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

.PHONY: builddir
builddir:
	@mkdir -p build

.PHONY: clean
clean: builddeps
	@echo " $(GRAY)🛁 Cleaning$(RESET)"
	@rm -frv $(BUILD_DIR)

.PHONY: check
check: builddeps
	@echo " $(PINK)👮 Checking$(RESET)"
	$(GOLINT) -set_exit_status ./...

.PHONY: test
test: builddir check
	@echo " $(GREEN)✔️ Testing$(RESET)"
	$(RICHGO) test -covermode=count -coverprofile=build/coverage.out ./...

.PHONY: binary
binary: builddir test
	@echo " $(BLUE)🔨 Building$(RESET)"
	$(RICHGO) build -o $(BIN) $(CLI_DIR)

.PHONY: run
run: builddir binary
	@echo " $(RED)🏃 Running$(RESET)"
	cd $(BUILD_DIR) && $(BIN) $(args)
