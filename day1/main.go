package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numbersInString = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
}

var re *regexp.Regexp = regexp.MustCompile("[0-9]")

// var reNumbersString = regexp.MustCompile("(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)|[0-9]")

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	counter := 0

	for scanner.Scan() {
		// fmt.Printf("reading this line: %s\n", scanner.Text())
		val := find2Number(scanner.Text())
		// fmt.Printf("value found: %d\n", val)
		counter += val
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("final value: %d\n", counter)

}

// func find2Digit(line string) int {
// 	allDigit := re.FindAllString(line, -1)
// 	val, err := strconv.Atoi(allDigit[0] + allDigit[len(allDigit)-1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return val

// }

func find2Number(line string) int {

	var firstDigit, lastDigit = "", ""
	var firstIdx, lastIdx = len(line), 0

	for numberWord := range numbersInString {
		startIdx := strings.Index(line, numberWord)
		if startIdx != -1 {
			if firstDigit == "" {
				//first found ever
				firstDigit = numbersInString[numberWord]
				lastDigit = numbersInString[numberWord]
				firstIdx = startIdx
				lastIdx = startIdx
				continue
			}
			if startIdx < firstIdx {
				firstIdx = startIdx
				firstDigit = numbersInString[numberWord]
			}
			if startIdx > lastIdx {
				lastIdx = startIdx
				lastDigit = numbersInString[numberWord]
			}

		}
	}

	// for numeric number
	// allDigit := re.FindAllString(line, -1)

	// var firstNumberIdx, lastNumberIdx int

	// if len(allDigit) == 1 {
	// 	firstNumberIdx = strings.Index(line, allDigit[0])
	// 	lastNumberIdx = firstNumberIdx
	// 	if firstNumberIdx < firstIdx {
	// 		firstDigit = allDigit[0]
	// 	}

	// 	if lastNumberIdx > lastIdx {
	// 		lastDigit = allDigit[0]
	// 	}

	// } else if len(allDigit) != 0 {

	// 	firstNumberIdx = strings.Index(line, allDigit[0])
	// 	lastNumberIdx = strings.Index(line, allDigit[1])

	// 	if firstNumberIdx < firstIdx {
	// 		firstDigit = allDigit[0]
	// 	}

	// 	if lastNumberIdx > lastIdx {
	// 		lastDigit = allDigit[1]
	// 	}
	// }

	val, err := strconv.Atoi(firstDigit + lastDigit)

	if err != nil {
		log.Fatal(err)
	}

	return val

}
