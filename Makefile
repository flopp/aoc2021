.PHONY: format
format:
	@go fmt day01/*.go
	@go fmt day02/*.go
	@go fmt day03/*.go
	@go fmt day04/*.go

.PHONY: 01
01:
	@echo "expected: 7"
	@go run day01/main.go 1 < day01/test.txt
	@go run day01/main.go 1 < day01/puzzle.txt
	@echo "expected: 5"
	@go run day01/main.go 3 < day01/test.txt
	@go run day01/main.go 3 < day01/puzzle.txt

.PHONY: 02
02:
	@echo "expected: 150"
	@go run day02/main.go part1 < day02/test.txt
	@go run day02/main.go part1 < day02/puzzle.txt
	@echo "expected: 900"
	@go run day02/main.go part2 < day02/test.txt
	@go run day02/main.go part2 < day02/puzzle.txt

.PHONY: 03
03:
	@echo "expected: 198"
	@go run day03/main.go part1 < day03/test.txt
	@go run day03/main.go part1 < day03/puzzle.txt
	@echo "expected: 230"
	@go run day03/main.go part2 < day03/test.txt
	@go run day03/main.go part2 < day03/puzzle.txt

.PHONY: 04
04:
	@echo "expected: 4512"
	@go run day04/main.go part1 < day04/test.txt
	@go run day04/main.go part1 < day04/puzzle.txt
	@echo "expected: 1924"
	@go run day04/main.go part2 < day04/test.txt
	@go run day04/main.go part2 < day04/puzzle.txt