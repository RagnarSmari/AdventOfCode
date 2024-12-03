package DayThree

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func readData() ([]string, error) {
	file, err := os.Open("./DayThree/data.txt")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Unable to close file")
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

type mule struct {
	first  int
	second int
}

func findUncorruptedMulesInLine(line string) ([]mule, error) {
	var mules []mule
	isEnabled := true
	// What do we need to get a correct mule? Need
	for index := 0; index < len(line); index++ {

		if line[index] == 'd' && line[index+1] == 'o' && line[index+2] == 'n' && line[index+3] == '\'' && line[index+4] == 't' && line[index+5] == '(' && line[index+6] == ')' {
			isEnabled = false
		}

		if line[index] == 'd' && line[index+1] == 'o' && line[index+2] == '(' && line[index+3] == ')' {
			isEnabled = true
		}

		if line[index] == 'm' && index+3 < len(line) {
			if line[index+1] == 'u' && line[index+2] == 'l' && line[index+3] == '(' {
				// We got ourselves a uncorrupted mule.
				firstNumber := ""
				secondNumber := ""
				isOnFirstNumber := true
				isCorrupted := false

				for i := index + 4; i < len(line); i++ {
					// Even in the mul instruction we can have corrupted character.
					// If the character is neither a closing ) or , it should be a number

					if line[i] == ',' { // Begin on the second number
						isOnFirstNumber = false
						continue
					}

					if line[i] == 41 { // ')'
						break
					}
					// At this point the character is neither ) or , so it should be a number
					if unicode.IsDigit(rune(line[i])) {
						number := int(line[i] - '0')
						if isOnFirstNumber {
							firstNumber += strconv.Itoa(number)
						} else {
							secondNumber += strconv.Itoa(number)
						}
					} else {
						isCorrupted = true
						break
					}
				}

				if isCorrupted {
					continue
				}
				first, err1 := strconv.Atoi(firstNumber)
				second, err2 := strconv.Atoi(secondNumber)

				if err1 == nil && err2 == nil && isEnabled {
					mules = append(mules, mule{first, second})
				} else {
					continue
				}
			}
		} else {
			continue
		}

	}
	return mules, nil
}

func DayThree() {
	data, err := readData()
	if err != nil {
		log.Fatal("Unable to read data")
	}

	total := 0
	dataWithoutNewLines := ""
	for _, line := range data {
		dataWithoutNewLines += line
	}
	uncorruptedMules, err := findUncorruptedMulesInLine(dataWithoutNewLines)
	for _, mule := range uncorruptedMules {
		totalMule := mule.first * mule.second
		total += totalMule
	}
	fmt.Printf(strconv.Itoa(total))
}
