package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const RequiredIngredients int = 100

type Ingredient struct {
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func NewIngredient(line string) *Ingredient {
	var newIngredient Ingredient

	re := regexp.MustCompile(`(\w+) (-?\d+)`)
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		if len(match) != 3 {
			continue
		}
		value, _ := strconv.Atoi(match[2])
		switch match[1] {
		case "capacity":
			newIngredient.capacity = value
		case "durability":
			newIngredient.durability = value
		case "flavor":
			newIngredient.flavor = value
		case "texture":
			newIngredient.texture = value
		case "calories":
			newIngredient.calories = value
		}
	}
	return &newIngredient
}

func getMaximumCookieScore(position, endPos, start int, current []int, getScore func([]int) int) int {
	if position == endPos {
		if start != RequiredIngredients {
			return 0
		}
		return getScore(current)
	}

	maxScore := 0
	for i := 0; i <= RequiredIngredients-start; i++ {
		current[position] = i
		score := getMaximumCookieScore(position+1, endPos, start+i, current, getScore)
		maxScore = max(maxScore, score)
	}
	return maxScore
}

func getCookieScoreFunction(ingredients []*Ingredient, calorieLimit int) func([]int) int {
	return func(nums []int) int {
		var capacityScore, durabilityScore, flavorScore, textureScore, calories int
		for idx := range nums {
			capacityScore += nums[idx] * ingredients[idx].capacity
			durabilityScore += nums[idx] * ingredients[idx].durability
			flavorScore += nums[idx] * ingredients[idx].flavor
			textureScore += nums[idx] * ingredients[idx].texture
			calories += nums[idx] * ingredients[idx].calories
		}

		if calorieLimit != 0 && calories != calorieLimit {
			return 0
		}

		capacityScore = max(0, capacityScore)
		durabilityScore = max(0, durabilityScore)
		textureScore = max(0, textureScore)
		flavorScore = max(0, flavorScore)

		return capacityScore * durabilityScore * textureScore * flavorScore
	}
}

func part1(ingredientsRecipe []*Ingredient) {
	count := make([]int, len(ingredientsRecipe))
	scoreFunction := getCookieScoreFunction(ingredientsRecipe, 0)

	highestScore := getMaximumCookieScore(0, len(ingredientsRecipe), 0, count, scoreFunction)

	fmt.Printf("Part 1: %d\n", highestScore)
}

func part2(ingredientsRecipe []*Ingredient) {
	count := make([]int, len(ingredientsRecipe))
	const ExpectedCalories int = 500
	scoreFunction := getCookieScoreFunction(ingredientsRecipe, ExpectedCalories)

	highestScore := getMaximumCookieScore(0, len(ingredientsRecipe), 0, count, scoreFunction)

	fmt.Printf("Part 2: %d\n", highestScore)
}

func parseInput(input []string) []*Ingredient {
	var result []*Ingredient
	for _, line := range input {
		currentIngredient := NewIngredient(line)
		if currentIngredient != nil {
			result = append(result, currentIngredient)
		}
	}
	return result
}

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("File name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	fileContent, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	input := strings.Split(string(fileContent), "\n")

	parsedInput := parseInput(input)

	part1(parsedInput)
	part2(parsedInput)
}
