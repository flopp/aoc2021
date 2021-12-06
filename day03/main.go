package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1 := true
	switch {
	case os.Args[1] == "part1":
		part1 = true
	case os.Args[1] == "part2":
		part1 = false
	default:
		panic(fmt.Errorf("bad part option: <%s>", os.Args[1]))
	}

	diagnostic_report := make([][]bool, 0)
	line_length := -1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		report_line := make([]bool, 0)
		for _, bit := range line {
			switch {
			case bit == '1':
				report_line = append(report_line, true)
			case bit == '0':
				report_line = append(report_line, false)
			default:
				panic(fmt.Errorf("bad line: <%s>", line))
			}
		}

		if line_length < 0 {
			line_length = len(report_line)
		} else {
			if len(report_line) != line_length {
				panic(fmt.Errorf("bad line: <%s>", line))
			}
		}
		diagnostic_report = append(diagnostic_report, report_line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if part1 {
		count1 := make([]int, len(diagnostic_report[0]))

		for _, report_line := range diagnostic_report {
			for index, bit := range report_line {
				if bit {
					count1[index] += 1
				}
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}

		gamma_rate := 0
		epsilon_rate := 0
		for _, count := range count1 {
			if count*2 > len(diagnostic_report) {
				// 1 is most common
				gamma_rate = (gamma_rate << 1) + 1
				epsilon_rate = (epsilon_rate << 1)
			} else {
				// 0 is most common
				gamma_rate = (gamma_rate << 1)
				epsilon_rate = (epsilon_rate << 1) + 1
			}
		}

		power_consumption := gamma_rate * epsilon_rate

		fmt.Printf("%d\n", power_consumption)
	} else {
		// oxygen generator rating
		filtered_diagnostic_report := append([][]bool{}, diagnostic_report...)
		for index := 0; index < line_length; index += 1 {
			if len(filtered_diagnostic_report) == 1 {
				break
			}
			count1 := 0
			for _, report_line := range filtered_diagnostic_report {
				if report_line[index] {
					count1 += 1
				}
			}
			filtered2 := make([][]bool, 0, len(filtered_diagnostic_report))
			if 2*count1 >= len(filtered_diagnostic_report) {
				for _, report_line := range filtered_diagnostic_report {
					if report_line[index] {
						filtered2 = append(filtered2, report_line)
					}
				}
			} else {
				for _, report_line := range filtered_diagnostic_report {
					if !report_line[index] {
						filtered2 = append(filtered2, report_line)
					}
				}
			}
			filtered_diagnostic_report = filtered2
		}
		if len(filtered_diagnostic_report) != 1 {
			panic(fmt.Errorf("len == %d", len(filtered_diagnostic_report)))
		}
		oxygen_generator_rating := 0
		for _, bit := range filtered_diagnostic_report[0] {
			oxygen_generator_rating = (oxygen_generator_rating << 1)
			if bit {
				oxygen_generator_rating += 1
			}
		}

		// CO2 scrubber rating
		filtered_diagnostic_report = append([][]bool{}, diagnostic_report...)
		for index := 0; index < line_length; index += 1 {
			if len(filtered_diagnostic_report) == 1 {
				break
			}
			count1 := 0
			for _, report_line := range filtered_diagnostic_report {
				if report_line[index] {
					count1 += 1
				}
			}
			filtered2 := make([][]bool, 0, len(filtered_diagnostic_report))
			if 2*count1 >= len(filtered_diagnostic_report) {
				// 1 is most common, or 0 and 1 are equally common => take 0 lines
				for _, report_line := range filtered_diagnostic_report {
					if !report_line[index] {
						filtered2 = append(filtered2, report_line)
					}
				}
			} else {
				// 0 is most common => take 1 lines
				for _, report_line := range filtered_diagnostic_report {
					if report_line[index] {
						filtered2 = append(filtered2, report_line)
					}
				}
			}
			filtered_diagnostic_report = filtered2
		}
		if len(filtered_diagnostic_report) != 1 {
			panic(fmt.Errorf("len == %d", len(filtered_diagnostic_report)))
		}
		co2_scrubber_rating := 0
		for _, bit := range filtered_diagnostic_report[0] {
			co2_scrubber_rating = (co2_scrubber_rating << 1)
			if bit {
				co2_scrubber_rating += 1
			}
		}

		life_support_rating := oxygen_generator_rating * co2_scrubber_rating
		fmt.Printf("%d\n", life_support_rating)
	}
}