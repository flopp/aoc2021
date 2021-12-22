define run1
	@go run day$@/main.go part1 < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part1 < day$@/puzzle.txt
	@echo
endef

define run2
	@go run day$@/main.go part2 < day$@/test.txt
	@echo "=>"
	@go run day$@/main.go part2 < day$@/puzzle.txt
	@echo
endef

.PHONY: all
all: 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20

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
	go fmt day13/main.go
	go fmt day14/main.go
	go fmt day15/main.go
	go fmt day16/main.go
	go fmt day17/main.go
	go fmt day18/main.go
	go fmt day19/main.go
	go fmt day20/main.go

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
	@echo "=>"
	@go run day$@/main.go part1 < day$@/puzzle.txt
	@echo
	@echo "expected: 36"
	@go run day$@/main.go part2 < day$@/test1.txt
	@echo "expected: 103"
	@go run day$@/main.go part2 < day$@/test2.txt
	@echo "expected: 3509"
	@go run day$@/main.go part2 < day$@/test3.txt
	@echo "=>"
	@go run day$@/main.go part2 < day$@/puzzle.txt
	@echo

.PHONY: 13
13:
	@echo "expected: 17"
	$(run1)
	@echo "expected: 0"
	$(run2)

.PHONY: 14
14:
	@echo "expected: 1588"
	$(run1)
	@echo "expected: 2188189693529"
	$(run2)

.PHONY: 15
15:
	@echo "expected: 40"
	$(run1)
	@echo "expected: 315"
	$(run2)

.PHONY: 16
16:
	@echo "expected: 16 23 31"
	@go run day$@/main.go part1 < day$@/test1a.txt
	@go run day$@/main.go part1 < day$@/test1b.txt
	@go run day$@/main.go part1 < day$@/test1c.txt
	@go run day$@/main.go part1 < day$@/test1d.txt
	@echo "=>"
	@go run day$@/main.go part1 < day$@/puzzle.txt
	@echo
	@echo "expected: 3 54 7 9 1 0 0 1"
	@go run day$@/main.go part2 < day$@/test2a.txt
	@go run day$@/main.go part2 < day$@/test2b.txt
	@go run day$@/main.go part2 < day$@/test2c.txt
	@go run day$@/main.go part2 < day$@/test2d.txt
	@go run day$@/main.go part2 < day$@/test2e.txt
	@go run day$@/main.go part2 < day$@/test2f.txt
	@go run day$@/main.go part2 < day$@/test2g.txt
	@go run day$@/main.go part2 < day$@/test2h.txt
	@echo "=>"
	@go run day$@/main.go part2 < day$@/puzzle.txt
	@echo

.PHONY: 17
17:
	@echo "expected: 45"
	$(run1)
	@echo "expected: 112"
	$(run2)

.PHONY: 18
18:
	@echo "expected: 4140"
	$(run1)
	@echo "expected: 3993"
	$(run2)

.PHONY: 19
19:
	@echo "expected: 79"
	$(run1)
	@echo "expected: ?"
	$(run2)

.PHONY: 20
20:
	@echo "expected: 35"
	$(run1)
	@echo "expected: 3351"
	$(run2)

.PHONY: 21
21:
	@echo "expected: 739785"
	$(run1)
	@echo "expected: ?"
	$(run2)
