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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isRowSafe(row []int) (bool, int){
	var isIncrementing bool
	var isDecrementing bool
	var totalUnsafeNumbers = 0;
	var isRowSafe = true;

	for index, value := range row {
		if index == 0 {
			continue
		}

		var oldValue = row[index-1]
		var diff = abs(value - oldValue)

		if diff < 1 || diff > 3 {
			totalUnsafeNumbers++
			isRowSafe = false
		}

		if value > oldValue {
			isIncrementing = true
		}

		if value < oldValue {
			isDecrementing = true
		}

		if isIncrementing && isDecrementing {
			totalUnsafeNumbers++
			isRowSafe = false
		}
	}

	return isRowSafe, totalUnsafeNumbers)

}


func canElementBeRemovedToMakeRowSafe(row []int) bool {
	for index, value := range row {

	}
}

func DayTwo() {
	fmt.Printf("Day two\n")
	records := readCsvFile("data.csv")
	var recordsInInt [][]int

	for _, record := range records {
		var parsedList = parseListToListOfNumber(strings.Split(record[0], " "))
		recordsInInt = append(recordsInInt, parsedList)
	}

	var totalSafe = 0
	for _, record := range recordsInInt {
		isRowSafeBool, totalUnsafeNumbers := isRowSafe(record)
		if isRowSafeBool {
			totalSafe++;
		}
		if totalUnsafeNumbers == 1 {

		}
	}
	fmt.Printf("Part one: " + strconv.Itoa(totalSafe) + "\n")
}
