define run1
	@go run day$@/main.go part1 < day$@/test.txt
	@echo "part 1 =>"
	@go run day$@/main.go part1 < day$@/puzzle.txt
endef

define run2
	@go run day$@/main.go part2 < day$@/test.txt
	@echo "part 2 =>"
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
	go fmt day07/main.go
	go fmt day08/main.go
	go fmt day09/main.go
	go fmt day10/main.go
	go fmt day11/main.go
	go fmt day12/main.go

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

.PHONY: 07
07:
	@echo "expected: 37"
	$(run1)
	@echo "expected: 168"
	$(run2)

.PHONY: 08
08:
	@echo "expected: 26"
	$(run1)
	@echo "expected: 61229"
	$(run2)

.PHONY: 09
09:
	@echo "expected: 15"
	$(run1)
	@echo "expected: 1134"
	$(run2)

.PHONY: 10
10:
	@echo "expected: 26397"
	$(run1)
	@echo "expected: 288957"
	$(run2)

.PHONY: 11
11:
	@echo "expected: 1656"
	$(run1)
	@echo "expected: 195"
	$(run2)

.PHONY: 12
12:
	@echo "expected: 10"
	@go run day$@/main.go part1 < day$@/test1.txt
	@echo "expected: 19"
	@go run day$@/main.go part1 < day$@/test2.txt
	@echo "expected: 226"
	@go run day$@/main.go part1 < day$@/test3.txt
	@echo "part 1 =>"
	@go run day$@/main.go part1 < day$@/puzzle.txt
	@echo "expected: 36"
	@go run day$@/main.go part2 < day$@/test1.txt
	@echo "expected: 103"
	@go run day$@/main.go part2 < day$@/test2.txt
	@echo "expected: 3509"
	@go run day$@/main.go part2 < day$@/test3.txt
	@echo "part 2 =>"
	@go run day$@/main.go part2 < day$@/puzzle.txt