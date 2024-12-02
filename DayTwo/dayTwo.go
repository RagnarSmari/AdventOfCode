package DayTwo

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open("./DayTwo/data.csv")
	if err != nil {
		log.Fatal("Unable to read input file " + filePath)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Unable to close file " + filePath)
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV file")
	}
	return records
}

func parseListToListOfNumber(list []string) []int {
	var result []int
	for _, item := range list {
		value, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal("Unable to convert item to int")
		}
		result = append(result, value)
	}
	return result
}

func isSequenceSafe(row []int) bool {
	if len(row) < 2 {
		return true
	}

	isDecreasing := row[1] < row[0]

	for index := 1; index < len(row); index++ {
		oldValue, value := row[index-1], row[index]
		if isDecreasing {
			if value < oldValue-3 || value >= oldValue {
				return false
			}
		} else {
			if value > oldValue+3 || value <= oldValue {
				return false
			}
		}
	}
	return true
}

func isRowSafe(row []int) bool {

	if isSequenceSafe(row) {
		return true
	}

	for index := 0; index < len(row); index++ {
		if index > 0 && row[index] == row[index-1] {
			continue
		}

		modifiedRow := append([]int{}, row[:index]...)
		modifiedRow = append(modifiedRow, row[index+1:]...)

		if isSequenceSafe(modifiedRow) {
			return true
		}
	}
	return false
}

func DayTwo() {
	fmt.Printf("Day two\n")
	records := readCsvFile("data.csv")
	var recordsInInt [][]int

	for _, record := range records {
		var parsedList = parseListToListOfNumber(strings.Split(record[0], " "))
		recordsInInt = append(recordsInInt, parsedList)
	}

	totalSafe := 0
	for _, record := range recordsInInt {
		if isRowSafe(record) {
			totalSafe += 1
		}
	}
	fmt.Printf("Part one: " + strconv.Itoa(totalSafe) + "\n")
}
