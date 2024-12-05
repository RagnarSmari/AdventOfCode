package DayFive

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]string, []string) {
	file, err := os.Open("./DayFive/data.txt")
	if err != nil {
		log.Fatal("Unable to open file")
	}

	var pageRules, pages []string
	var isPages bool

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isPages = true
			continue
		}

		if isPages {
			pages = append(pages, line)
		} else {
			pageRules = append(pageRules, line)
		}
	}

	return pageRules, pages
}

func rulesToDict(rules []string) map[int][]int {
	var result = make(map[int][]int)

	for _, rule := range rules {
		numbers := strings.Split(rule, "|")
		firstNumber, _ := strconv.Atoi(numbers[0])
		secondNumber, _ := strconv.Atoi(numbers[1])
		result[firstNumber] = append(result[firstNumber], secondNumber)
	}
	return result
}

func processPages(pages []string) [][]int {
	var result [][]int
	for _, page := range pages {
		var tempNumbers []int
		numbers := strings.Split(page, ",")
		for _, number := range numbers {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal("Unable to convert number to int")
			}
			tempNumbers = append(tempNumbers, numberInt)
		}
		result = append(result, tempNumbers)
	}
	return result
}

func findNumberInList(list []int, number int) int {
	for index, item := range list {
		if item == number {
			return index
		}
	}
	return -1
}

func getGetOrderedAndUnorderedPages(pages [][]int, rules map[int][]int) ([][]int, [][]int) {
	var ordered [][]int
	var unOrdered [][]int

	for _, page := range pages {
		isCorrectlyOrdered := true
		for index, number := range page {
			// Loop through each numbers rules and check if the pages after are in the rule
			numbersRules := rules[number]
			for _, ruleNumber := range numbersRules {

				ruleNumberIndexInPage := findNumberInList(page, ruleNumber)
				if ruleNumberIndexInPage < index && ruleNumberIndexInPage != -1 { // We did not find the number and the number should
					isCorrectlyOrdered = false
					break
				}
			}
			if !isCorrectlyOrdered {
				break
			}
		}
		if isCorrectlyOrdered {
			ordered = append(ordered, page)
		} else {
			unOrdered = append(unOrdered, page)
		}
	}
	return ordered, unOrdered
}

func calculateAndSumMiddlePages(pages [][]int) int {
	total := 0
	for _, page := range pages {
		midIndex := len(page) / 2
		middleNumber := page[midIndex]
		total += middleNumber
	}
	return total
}

func orderPages(pages [][]int, rules map[int][]int) [][]int {
	// Order a page to the correct order based on its rules.
	// Each number has a rule and a rule means that the numbers in the pages rules can come after them. Not before
	var ordered [][]int
	for _, page := range pages {
		ordered = append(ordered, orderPage(page, rules))
	}
	return ordered
}

func relocateElementToIndex(page []int, currentIndex int, targetIndex int) []int {
	if currentIndex == targetIndex || currentIndex >= len(page) || targetIndex >= len(page) {
		return page
	}

	element := page[currentIndex]
	page = append(page[:currentIndex], page[currentIndex+1:]...)

	if targetIndex > currentIndex {
		targetIndex--
	}
	page = append(page[:targetIndex], append([]int{element}, page[targetIndex:]...)...)

	return page

}

func orderPage(page []int, rules map[int][]int) []int {
	if len(page) == 1 {
		return page
	}
	for index, number := range page {
		numbersRules := rules[number]
		for _, ruleNumber := range numbersRules {
			ruleNumberIndexInPage := findNumberInList(page, ruleNumber)
			if ruleNumberIndexInPage < index && ruleNumberIndexInPage != -1 { // The number is in the incorrect place

				page = relocateElementToIndex(page, index, ruleNumberIndexInPage)
				return orderPage(page, rules)
			}
		}
	}
	return page
}

func DayFive() {

	rules, pages := readInput()
	processedRules := rulesToDict(rules)
	processedPages := processPages(pages)
	orderedPages, unOrderedPages := getGetOrderedAndUnorderedPages(processedPages, processedRules)
	total := calculateAndSumMiddlePages(orderedPages)
	//for _, page := range unOrderedPages {
	//	fmt.Println(page)
	//}
	//for _, page := range orderings {
	//	fmt.Println(page)
	//}

	unOrderedTotal := calculateAndSumMiddlePages(orderPages(unOrderedPages, processedRules))

	// Calculate middle pages
	fmt.Println(total)
	fmt.Println(unOrderedTotal)
}
