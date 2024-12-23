package lib

import (
	"bufio"
	"os"
)

func ReadFileByLine(filename string, callback func(string)) {
	file, _ := os.Open(filename)
	defer file.Close()

	r := bufio.NewScanner(file)

	for r.Scan() {
		line := r.Text()
		callback(line)
	}

	file.Close()
}
