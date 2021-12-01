SHELL=/bin/bash

YEAR ?= $(shell /bin/date +"%Y")
DAY ?= $(shell /bin/date +"%d")

AOC_SESSION_COOKIE ?= $(shell cat ./aoc-session-cookie)
AOC_INPUT_FILE = https://adventofcode.com/$(YEAR)/day/$(shell echo $(DAY) | sed 's/^0*//')/input

CURRENTDAY = $(YEAR)/day$(DAY)

all: $(CURRENTDAY)

$(CURRENTDAY):
	@mkdir -p ./$(YEAR)/
	@cp -r template ./$(CURRENTDAY)
	@echo "folder ./$(CURRENTDAY) from template created"
	@curl --cookie $(AOC_SESSION_COOKIE) $(AOC_INPUT_FILE) -o ./$(CURRENTDAY)/input.txt
