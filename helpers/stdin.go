package helpers

import (
	"bufio"
	"os"
)

type LineHandler func(string)

func ReadStdin(handler LineHandler) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		handler(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
