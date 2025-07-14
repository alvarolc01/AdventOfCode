package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	BOARD_SIDE int = 5
)

type Position struct {
	row, col int
}
type Board struct {
	nums  map[int]Position
	found [][]bool
}

func NewBoard(boardLines []string) (*Board, error) {
	found := make([][]bool, BOARD_SIDE)
	for i := range BOARD_SIDE {
		found[i] = make([]bool, BOARD_SIDE)
	}

	nums := make(map[int]Position)
	for row, line := range boardLines {
		line = strings.TrimSpace(line)
		numsLine := strings.Fields(line)

		for col, numStr := range numsLine {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			nums[num] = Position{
				row: row,
				col: col,
			}
		}
	}

	return &Board{
		nums:  nums,
		found: found,
	}, nil
}

func (b *Board) PlayNum(num int) {
	if pos, ok := b.nums[num]; ok {
		delete(b.nums, num)
		b.found[pos.row][pos.col] = true
	}
}

func (b *Board) IsWon() bool {
	for _, row := range b.found {
		rowWon := true
		for _, foundNum := range row {
			if !foundNum {
				rowWon = false
				break
			}
		}

		if rowWon {
			return true
		}
	}

	for col := range BOARD_SIDE {
		colWon := true
		for row := range BOARD_SIDE {
			if !b.found[row][col] {
				colWon = false
				break
			}
		}

		if colWon {
			return true
		}
	}

	return false
}

func (b *Board) ScoreWithTurn(lastPlayedValue int) int {
	sumUnusedNums := 0
	for key := range b.nums {
		sumUnusedNums += key
	}

	return sumUnusedNums * lastPlayedValue
}

func part1(nums []int, boards []*Board) {
	score := -1

	for i := 0; i < len(nums) && score == -1; i++ {
		for _, board := range boards {
			board.PlayNum(nums[i])
			if board.IsWon() {
				score = board.ScoreWithTurn(nums[i])
				break
			}
		}
	}

	fmt.Printf("Part 1: %d\n", score)
}

func part2(nums []int, boards []*Board) {
	score := 0

	for i := 0; i < len(nums) && len(boards) > 0; i++ {
		boardsNextTurn := make([]*Board, 0)
		for _, board := range boards {
			board.PlayNum(nums[i])
			if board.IsWon() {
				score = board.ScoreWithTurn(nums[i])
			} else {
				boardsNextTurn = append(boardsNextTurn, board)
			}
		}
		boards = boardsNextTurn
	}

	fmt.Printf("Part 2: %d\n", score)

}

func parseInput(blocks []string) ([]int, []*Board, error) {
	bingoNums := strings.Split(blocks[0], ",")
	nums := make([]int, len(bingoNums))
	for i, num := range bingoNums {
		convertedNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, nil, err
		}
		nums[i] = convertedNum
	}

	boards := make([]*Board, len(blocks[1:]))
	for idx, block := range blocks[1:] {
		linesBlock := strings.Split(block, "\n")
		board, err := NewBoard(linesBlock)
		if err != nil {
			return nil, nil, err
		}

		boards[idx] = board
	}

	return nums, boards, nil
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n\n")

	nums, boards, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}

	part1(nums, boards)

	_, boardsPart2, err := parseInput(input)
	if err != nil {
		fmt.Println("error parsing input:", err)
		os.Exit(1)
	}
	part2(nums, boardsPart2)
}
