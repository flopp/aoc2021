define run1
	@go run day$@/main.go part1 < day$@/test.txt
	@go run day$@/main.go part1 < day$@/puzzle.txt
endef

define run2
	@go run day$@/main.go part2 < day$@/test.txt
	@go run day$@/main.go part2 < day$@/puzzle.txt
endef

.PHONY: format
format:
	go fmt helpers/*.go
	go fmt day01/main.go
	go fmt day02/main.go
	go fmt day03/main.go
	go fmt day04/main.go
	go fmt day05/main.go
	go fmt day06/main.go

.PHONY: 01
01:
	@echo "expected: 7"
	$(run1)
	@echo "expected: 5"
	$(run2)

.PHONY: 02
02:
	@echo "expected: 150"
	$(run1)
	@echo "expected: 900"
	$(run2)

.PHONY: 03
03:
	@echo "expected: 198"
	$(run1)
	@echo "expected: 230"
	$(run2)

.PHONY: 04
04:
	@echo "expected: 4512"
	$(run1)
	@echo "expected: 1924"
	$(run2)

.PHONY: 05
05:
	@echo "expected: 5"
	$(run1)
	@echo "expected: 12"
	$(run2)

.PHONY: 06
06:
	@echo "expected: 5934"
	$(run1)
	@echo "expected: 26984457539"
	$(run2)
