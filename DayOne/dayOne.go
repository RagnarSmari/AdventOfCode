package DayOne

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open("./DayOne/data.csv")
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
	csvReader.Comma = ' '
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse CSV file")
	}
	return records
}

func processInput(records [][]string) ([]int, []int) {
	var listOne []int
	var listTwo []int

	for _, record := range records {
		// Trim space and convert to integer
		var itemOne = record[0]
		firstItem, err := strconv.Atoi(strings.TrimSpace(itemOne))
		if err != nil {
			log.Fatalf("Unable to convert first item '%s' to int", record[0])
		}
		var itemTwo = record[3]
		secondItem, err := strconv.Atoi(strings.TrimSpace(itemTwo))
		if err != nil {
			log.Fatalf("Unable to convert second item '%s' to int", record[1])
		}
		listOne = append(listOne, firstItem)
		listTwo = append(listTwo, secondItem)
	}
	return listOne, listTwo

}

func findDistanceBetweenTwoSortedLists(listOne []int, listTwo []int, total int) int {

	if len(listOne) == 0 || len(listTwo) == 0 {
		return total
	}

	distance := abs(listOne[0] - listTwo[0])

	total += distance

	return findDistanceBetweenTwoSortedLists(listOne[1:], listTwo[1:], total)
}

func findFrequencyOfNumberInList(number int, list []int, total int) int {
	if len(list) == 0 {
		return total
	}

	if list[0] == number {
		total++
	}
	return findFrequencyOfNumberInList(number, list[1:], total)
}

func findFrequencyInList(listOne []int, listTwo []int, total int, cacheMemory map[int]int) int {

	if len(listOne) == 0 {
		return total
	}

	val, ok := cacheMemory[listOne[0]]
	if ok {
		total += val
	} else {
		var frequency = findFrequencyOfNumberInList(listOne[0], listTwo, 0)
		var sum = listOne[0] * frequency
		cacheMemory[listOne[0]] = sum
		total += sum
	}
	return findFrequencyInList(listOne[1:], listTwo, total, cacheMemory)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func DayOne() {
	records := readCsvFile("data.csv")
	listOne, listTwo := processInput(records)
	sort.Ints(listOne)
	sort.Ints(listTwo)
	total := findDistanceBetweenTwoSortedLists(listOne, listTwo, 0)
	frequency := findFrequencyInList(listOne, listTwo, 0, make(map[int]int))
	fmt.Printf("Part one: " + strconv.Itoa(total) + "\n")
	fmt.Printf("Part two: " + strconv.Itoa(frequency) + "\n")
}
