package main

import (
	"fmt"

	"github.com/flopp/aoc2021/helpers"
)

func main() {
	diagnostic_report := make([][]bool, 0)
	line_length := -1
	helpers.ReadStdin(func(line string) {
		report_line := make([]bool, 0)
		for _, bit := range line {
			switch bit {
			case '1':
				report_line = append(report_line, true)
			case '0':
				report_line = append(report_line, false)
			default:
				panic("bad input")
			}
		}

		if line_length < 0 {
			line_length = len(report_line)
		} else {
			if len(report_line) != line_length {
				panic("bad input")
			}
		}
		diagnostic_report = append(diagnostic_report, report_line)
	})

	if helpers.Part1() {
		count1 := make([]int, len(diagnostic_report[0]))

		for _, report_line := range diagnostic_report {
			for index, bit := range report_line {
				if bit {
					count1[index]++
				}
			}
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

		fmt.Println(power_consumption)
	} else {
		// oxygen generator rating
		filtered_diagnostic_report := append([][]bool{}, diagnostic_report...)
		for index := 0; index < line_length; index++ {
			if len(filtered_diagnostic_report) == 1 {
				break
			}
			count1 := 0
			for _, report_line := range filtered_diagnostic_report {
				if report_line[index] {
					count1++
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
				oxygen_generator_rating++
			}
		}

		// CO2 scrubber rating
		filtered_diagnostic_report = append([][]bool{}, diagnostic_report...)
		for index := 0; index < line_length; index++ {
			if len(filtered_diagnostic_report) == 1 {
				break
			}
			count1 := 0
			for _, report_line := range filtered_diagnostic_report {
				if report_line[index] {
					count1++
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
				co2_scrubber_rating++
			}
		}

		life_support_rating := oxygen_generator_rating * co2_scrubber_rating
		fmt.Println(life_support_rating)
	}
}
