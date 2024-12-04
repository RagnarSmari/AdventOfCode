package DayFour

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readData() ([]string, error) {
	file, err := os.Open("./DayFour/data.txt")
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

type APos struct {
	line   int
	char   int
	isDone bool
}

func getListOfAs(data []string) []APos {
	lineLength := len(data[0]) - 1
	totalLines := len(data) - 1
	var listOfAs []APos

	for lineIndex, line := range data {
		for charIndex, char := range line {
			if char == 'M' {

				if lineIndex-2 >= 0 { // Can we go up?
					// Up and left
					if charIndex-2 >= 0 {
						if data[lineIndex-1][charIndex-1] == 'A' && data[lineIndex-2][charIndex-2] == 'S' {
							listOfAs = append(listOfAs, APos{lineIndex - 1, charIndex - 1, false})
						}
					}

					// Up and right
					if charIndex+2 <= lineLength {
						if data[lineIndex-1][charIndex+1] == 'A' && data[lineIndex-2][charIndex+2] == 'S' {
							listOfAs = append(listOfAs, APos{lineIndex - 1, charIndex + 1, false})
						}
					}
				}

				// Can we go down?
				if lineIndex+2 <= totalLines {
					if charIndex+2 <= lineLength { // Down right
						if data[lineIndex+1][charIndex+1] == 'A' && data[lineIndex+2][charIndex+2] == 'S' {
							listOfAs = append(listOfAs, APos{lineIndex + 1, charIndex + 1, false})
						}
					}

					if charIndex-2 >= 0 { // Down left
						if data[lineIndex+1][charIndex-1] == 'A' && data[lineIndex+2][charIndex-2] == 'S' {
							listOfAs = append(listOfAs, APos{lineIndex + 1, charIndex - 1, false})
						}
					}
				}
			}
		}
	}
	return listOfAs

}

func findXMAS(data []string) int {
	lineLength := len(data[0]) - 1
	totalLines := len(data) - 1

	total := 0
	for lineIndex, line := range data {

		for charIndex, char := range line {

			if char == 'X' { // Check around the char, both up and down

				if charIndex-3 >= 0 { // Backwards
					if line[charIndex-1] == 'M' && line[charIndex-2] == 'A' && line[charIndex-3] == 'S' {
						total++
					}
				}

				if charIndex+3 <= lineLength { // Frontwards
					if line[charIndex+1] == 'M' && line[charIndex+2] == 'A' && line[charIndex+3] == 'S' { // Just Front
						total++
					}
				}

				// Can we move the word length up?
				if lineIndex-3 >= 0 { // Up

					// Up
					if data[lineIndex-1][charIndex] == 'M' && data[lineIndex-2][charIndex] == 'A' && data[lineIndex-3][charIndex] == 'S' {
						total++
					}

					// Can we go left
					if charIndex-3 >= 0 { // Backwards
						// Up to left
						if data[lineIndex-1][charIndex-1] == 'M' && data[lineIndex-2][charIndex-2] == 'A' && data[lineIndex-3][charIndex-3] == 'S' {
							total++
						}
					}
					// Can we go right?
					// Up to right
					if charIndex+3 <= lineLength {
						if data[lineIndex-1][charIndex+1] == 'M' && data[lineIndex-2][charIndex+2] == 'A' && data[lineIndex-3][charIndex+3] == 'S' {
							total++
						}
					}
				}

				// Can me move the word length down?
				if lineIndex+3 <= totalLines {
					// Down
					if data[lineIndex+1][charIndex] == 'M' && data[lineIndex+2][charIndex] == 'A' && data[lineIndex+3][charIndex] == 'S' {
						total++
					}

					// Can we go right
					if charIndex+3 <= lineLength {
						// Down to right
						if data[lineIndex+1][charIndex+1] == 'M' && data[lineIndex+2][charIndex+2] == 'A' && data[lineIndex+3][charIndex+3] == 'S' {
							total++
						}
					}

					// Can we go left?
					// Down to left
					if charIndex-3 >= 0 {
						if data[lineIndex+1][charIndex-1] == 'M' && data[lineIndex+2][charIndex-2] == 'A' && data[lineIndex+3][charIndex-3] == 'S' {
							total++
						}
					}
				}
			}
		}
	}
	return total
}

func DayFour() {
	data, err := readData()
	if err != nil {
		log.Fatal("Unable to read data")
	}

	listOfAs := getListOfAs(data)
	total := 0
	for firstIndex, firstPos := range listOfAs {
		for secondIndex, secondAPos := range listOfAs {
			if firstIndex == secondIndex {
				continue
			}
			if firstPos.line == secondAPos.line && firstPos.char == secondAPos.char && !firstPos.isDone && !secondAPos.isDone {
				listOfAs[firstIndex].isDone = true
				listOfAs[secondIndex].isDone = true
				total += 1
			}
		}
	}
	fmt.Printf(strconv.Itoa(total))

}
