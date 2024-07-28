# Colors
GREEN=\033[0;32m
NC=\033[0m
YELLOW=\033[0;33m
BLUE=\033[0;34m

# Variables
TMP_DIR := ./tmp
STATIC_DIR := ./assets
BUILD_DIR := ./build
MAIN_FILE := ./main.go
TAILWIND_FLAGS := --minify
UTILS_DIR := ./utils
CSS_INPUT_FILE := ${UTILS_DIR}/input.css
CSS_OUTPUT_FILE := ${STATIC_DIR}/output.min.css
OUTPUT_FILE := ${BUILD_DIR}/main

.PHONY: live/tailwindcss live/templ live/build live/preview live clean run

default: ./.air.toml
	@air -c $<

live/tailwindcss: $(CSS_INPUT_FILE)
	@echo -e "${BLUE}Running tailwindcss...${NC}"
	@bunx tailwindcss -i $(CSS_INPUT_FILE) -o $(CSS_OUTPUT_FILE) $(TAILWIND_FLAGS)

live/templ:
	@echo -e "${BLUE}Generating templ code...${NC}"
	@templ generate --proxy="http://localhost:3000" --open-browser=false

live/build: ${MAIN_FILE}
	@echo -e "${BLUE}Building the go binary...${NC}"
	@go build -o ${OUTPUT_FILE}

live/preview:
	@echo -e "${GREEN}Running the binary...${NC}"
	@${OUTPUT_FILE}

live: live/tailwindcss live/templ live/build

run: ${MAIN_FILE}
	@echo -e "${GREEN}Running the main.go file...${NC}"
	@go $@ $<

clean:
	@echo -e "${YELLOW}Cleaning...${NC}"
	@rm -rf ${BUILD_DIR} ${TMP_DIR} ${CSS_OUTPUT_FILE} *_templ.go
