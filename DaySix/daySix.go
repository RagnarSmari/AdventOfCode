package DaySix

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type block struct {
	linePos int
	charPos int
}

type guard struct {
	xPos      int
	yPos      int
	direction rune
}

var originalPos = guard{
	0,
	0,
	'>',
}
var totalLineLength = 0

// Data - position of blocks e.g. '#', X pos of guard, Y pos of guard, error
func readInput() ([][]rune, guard) {
	file, err := os.Open("./DaySix/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var data [][]rune
	var blocks []block
	player := guard{
		xPos:      0,
		yPos:      0,
		direction: '<',
	}

	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		data = append(data, runes)
		totalLineLength = len(runes)
		for charIndex, char := range line {
			if char == '^' { // Found the starting pos of guard
				player.direction = char
				player.yPos = lineIndex
				player.xPos = charIndex
			}
			if char == '#' {
				blocks = append(blocks, block{lineIndex, charIndex})
			}
		}
		lineIndex++
	}
	return data, player
}

func calculateDistinctPositions(data [][]rune, player guard, total int) int {
	totalLines := len(data)
	// Check direction of guard
	if player.direction == '^' { // Up
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		length := 0
		for i := player.yPos; i >= 0; i-- {
			if data[i][player.xPos] == '.' {
				data[i][player.xPos] = 'X'
				length++
			}

			if data[i][player.xPos] == '#' {
				foundBlock = true
				total += length
				player.yPos = i + 1
				break
			}

		}
		if !foundBlock {
			total += length
			return total
		}
		// Turn the guard 90 degrees to the right
		player.direction = '>'
	} else if player.direction == 'v' { // Down
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		length := 0
		for i := player.yPos; i <= totalLineLength-1; i++ {

			if data[i][player.xPos] == '.' {
				length++
				data[i][player.xPos] = 'X'
			}

			if data[i][player.xPos] == '#' {
				foundBlock = true
				total += length
				player.yPos = i - 1
				break
			}

		}
		if !foundBlock {
			total += length
			for _, line := range data {
				fmt.Println(string(line))
			}
			return total
		}
		player.direction = '<'
	} else if player.direction == '<' { // left
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		length := 0
		for i := player.xPos; i >= 0; i-- {

			if data[player.yPos][i] == '.' {
				data[player.yPos][i] = 'X'
				length++
			}

			if data[player.yPos][i] == '#' {
				foundBlock = true
				total += length
				player.xPos = i + 1
				break
			}

		}
		if !foundBlock {
			total += length
			return total
		}
		player.direction = '^'
	} else if player.direction == '>' { // right
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		length := 0
		for i := player.xPos; i <= totalLines-1; i++ {

			if data[player.yPos][i] == '.' {
				data[player.yPos][i] = 'X'
				length++
			}

			if data[player.yPos][i] == '#' {
				foundBlock = true
				total += length
				player.xPos = i - 1
				// Mark all paths from the pos to this pos as X in the data
				break
			}

		}
		if !foundBlock {
			total += length
			return total
		}
		player.direction = 'v'
	}

	return calculateDistinctPositions(data, player, total)
}

// Try to traverse the grid, if we end up on the block in the params -- the guard is stuck in a loop and we return false as the guard could not traverse
func tryTraverse(data [][]rune, player guard, initialBlock block, totalCircles int, visited map[string]int) bool {
	totalLines := len(data)
	// Where is player?
	//fmt.Println("%d,%d,%c", player.yPos, player.xPos, player.direction)

	if player.direction == '^' { // Up
		foundBlock := false
		for i := player.yPos; i >= 0; i-- {

			if data[i][player.xPos] == '#' || (i == initialBlock.linePos && player.xPos == initialBlock.charPos) {
				// We found a initialBlock - check if we have already been in this initialBlock
				block := fmt.Sprintf("%d %d", i, player.xPos)
				val, exists := visited[block]
				if exists {
					visited[block] = val + 1
				} else {
					visited[block] = 1
				}
				if val >= 4 {
					return false
				}

				foundBlock = true
				player.yPos = i + 1
				break
			}

		}
		if !foundBlock {
			return true // Successfully traversed
		}
		// Turn the guard 90 degrees to the right
		player.direction = '>'
		return tryTraverse(data, player, initialBlock, totalCircles, visited)
	} else if player.direction == 'v' { // Down
		// Check if there is any initialBlock in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.yPos; i <= totalLines-1; i++ {

			if data[i][player.xPos] == '#' || (i == initialBlock.linePos && player.xPos == initialBlock.charPos) {
				// We found a initialBlock - check if we have already been in this initialBlock
				block := fmt.Sprintf("%d %d", i, player.xPos)
				val, exists := visited[block]
				if exists {
					visited[block] = val + 1
				} else {
					visited[block] = 1
				}
				if val >= 4 {
					return false
				}

				foundBlock = true
				player.yPos = i - 1
				break
			}

		}
		if !foundBlock {
			return true
		}
		player.direction = '<'
		return tryTraverse(data, player, initialBlock, totalCircles, visited)
	} else if player.direction == '<' { // left
		// Check if there is any initialBlock in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.xPos; i >= 0; i-- {

			if data[player.yPos][i] == '#' || (player.yPos == initialBlock.linePos && i == initialBlock.charPos) {
				block := fmt.Sprintf("%d %d", i, player.xPos)
				val, exists := visited[block]
				if exists {
					visited[block] = val + 1
				} else {
					visited[block] = 1
				}
				if val >= 4 {
					return false
				}

				foundBlock = true
				player.xPos = i + 1
				break
			}

		}
		if !foundBlock {
			return true
		}
		player.direction = '^'
		return tryTraverse(data, player, initialBlock, totalCircles, visited)
	} else if player.direction == '>' { // right
		// Check if there is any initialBlock in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.xPos; i <= totalLineLength-1; i++ {

			if data[player.yPos][i] == '#' || (player.yPos == initialBlock.linePos && i == initialBlock.charPos) {
				block := fmt.Sprintf("%d %d", i, player.xPos)
				val, exists := visited[block]
				if exists {
					visited[block] = val + 1
				} else {
					visited[block] = 1
				}
				if val >= 4 {
					return false
				}

				foundBlock = true
				player.xPos = i - 1
				break
			}

		}
		if !foundBlock {
			return true
		}
		player.direction = 'v'
		return tryTraverse(data, player, initialBlock, totalCircles, visited)
	}

	return tryTraverse(data, player, initialBlock, totalCircles, visited)

}

// Does not work sadly :'(
func findAllPossibleBoxPositions(data [][]rune, player guard, total int, triedBlocks map[string]bool) int {
	fmt.Printf("findAllPossibleBoxPositions called with player at (%d, %d) facing %c\n", player.yPos, player.xPos, player.direction)
	totalLines := len(data)
	// Check direction of guard
	if player.direction == '^' { // Up
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.yPos - 1; i >= 0; i-- {

			if i-1 < 0 {
				break
			}
			if data[i-1][player.xPos] != '#' {

				for y := player.xPos; y <= totalLineLength-1; y++ {
					if data[i][y] == '#' {
						triedBlock := fmt.Sprintf("%d %d", i-1, player.xPos)
						_, exists := triedBlocks[triedBlock]
						if !exists {
							couldTraverse := tryTraverse(data, originalPos, block{i - 1, player.xPos}, 0, make(map[string]int))
							if !couldTraverse {
								total++
							}
							triedBlocks[triedBlock] = true
						}
						break
					}
				}

			} else {
				// We found a box which is in front of us - aka we cant put a box there so we dont need to check
				foundBlock = true
				player.yPos = i
				player.direction = '>'
				break
			}

		}
		if !foundBlock {
			return total
		}
		return findAllPossibleBoxPositions(data, player, total, triedBlocks)

	} else if player.direction == 'v' { // Down
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.yPos + 1; i <= totalLines-1; i++ {

			if i+1 == totalLines {
				break
			}
			if i != totalLines && data[i+1][player.xPos] != '#' {

				for y := player.xPos; y >= 0; y-- { // Check on the left if there is a box
					if data[i][y] == '#' {
						triedBlock := fmt.Sprintf("%d %d", i+1, player.xPos)
						_, exists := triedBlocks[triedBlock]
						if !exists {
							couldTraverse := tryTraverse(data, originalPos, block{i + 1, player.xPos}, 0, make(map[string]int))
							if !couldTraverse {
								total++
							}
							triedBlocks[triedBlock] = true
						}
						break
					}
				}

			} else {
				foundBlock = true
				player.yPos = i
				player.direction = '<'
				break
			}

		}
		if !foundBlock {
			return total
		}
		return findAllPossibleBoxPositions(data, player, total, triedBlocks)
	} else if player.direction == '<' { // left
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.xPos - 1; i >= 0; i-- {
			if i-1 < 0 {
				break
			}

			if data[player.yPos][i-1] != '#' {
				for y := player.yPos; y >= 0; y-- {
					if data[y][i] == '#' {
						triedBlock := fmt.Sprintf("%d %d", player.yPos, i-1)
						_, exists := triedBlocks[triedBlock]
						if !exists {
							couldTraverse := tryTraverse(data, originalPos, block{player.yPos, i - 1}, 0, make(map[string]int))
							if !couldTraverse {
								total++
							}
							triedBlocks[triedBlock] = true
						}
						break
					}
				}
			} else {
				foundBlock = true
				player.xPos = i
				player.direction = '^'
				break

			}
		}
		if !foundBlock {
			return total
		}
		return findAllPossibleBoxPositions(data, player, total, triedBlocks)

	} else if player.direction == '>' { // right
		// Check if there is any block in the guard direction to the end of the map, if not, the game ended
		foundBlock := false
		for i := player.xPos + 1; i <= totalLineLength-1; i++ {
			if i+1 == totalLines {
				break
			}

			if data[player.yPos][i+1] != '#' {
				for y := player.yPos; y <= totalLines-1; y++ {
					if data[y][i] == '#' {
						triedBlock := fmt.Sprintf("%d %d", player.yPos, i+1)
						_, exists := triedBlocks[triedBlock]
						if !exists {
							couldTraverse := tryTraverse(data, originalPos, block{player.yPos, i + 1}, 0, make(map[string]int))
							if !couldTraverse {
								total++
							}
							triedBlocks[triedBlock] = true
						}
						break
					}
				}
			} else {
				foundBlock = true
				player.xPos = i
				player.direction = 'v'
				break
			}
		}
		if !foundBlock {
			return total
		}
		return findAllPossibleBoxPositions(data, player, total, triedBlocks)
	} else {
		return total
	}

}

func DaySix() {
	partTwoTotal := 0
	partOneTotal := 0
	data, player := readInput()
	originalPos = player
	for lineIndex, line := range data {
		for charIndex, char := range line {
			if char == '.' {
				canTraverse := tryTraverse(data, originalPos, block{linePos: lineIndex, charPos: charIndex}, 0, make(map[string]int))
				if !canTraverse {
					partTwoTotal++
				}
			}
		}
	}
	partOneTotal = calculateDistinctPositions(data, player, 0)
	//fmt.Println(originalPos.yPos, originalPos.xPos, originalPos.direction)
	//total := findAllPossibleBoxPositions(data, player, 0, make(map[string]bool)) --- DID NOT WORK
	fmt.Println(partOneTotal + 1) // Account for the initial position
	fmt.Println(partTwoTotal)
}
