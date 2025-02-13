package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	PlayerStartingHP int = 100
	MaximumRings     int = 2
)

type Character struct {
	hitPoints int
	damage    int
	armor     int
}

func (c *Character) AddEquipment(e *Equipment) {
	c.damage += e.damage
	c.armor += e.armor
}

func (c *Character) TurnsLasted(opponentDamage int) int {
	damagePerTurn := max(opponentDamage-c.armor, 1)
	return (c.hitPoints + damagePerTurn - 1) / damagePerTurn
}

type Equipment struct {
	cost   int
	damage int
	armor  int
}

func canDefeatBoss(boss *Character, selectedEquipment []Equipment) bool {
	player := &Character{hitPoints: PlayerStartingHP}
	for _, e := range selectedEquipment {
		player.AddEquipment(&e)
	}

	turnsPlayerLasts := player.TurnsLasted(boss.damage)
	turnsBossLasts := boss.TurnsLasted(player.damage)

	return turnsPlayerLasts >= turnsBossLasts
}

func backtrackingEquipmentCombinations(
	boss *Character,
	selectedEquipment []Equipment,
	numRings, posRing int,
	rings []Equipment,
	comparisonFunc func(int, int) int,
	costWinning bool,
) int {
	if numRings > MaximumRings {
		return math.MaxInt
	}
	currentCost := math.MaxInt
	wouldWin := canDefeatBoss(boss, selectedEquipment)

	if wouldWin == costWinning {
		totalCost := 0
		for _, eq := range selectedEquipment {
			totalCost += eq.cost
		}
		currentCost = totalCost
	}

	if len(rings) <= posRing {
		return currentCost
	}

	costSelectingCurrentRing := backtrackingEquipmentCombinations(
		boss,
		append(selectedEquipment, rings[posRing]),
		numRings+1, posRing+1,
		rings,
		comparisonFunc,
		costWinning,
	)

	costNotSelectingCurrentRing := backtrackingEquipmentCombinations(
		boss,
		selectedEquipment,
		numRings, posRing+1,
		rings,
		comparisonFunc,
		costWinning,
	)

	return comparisonFunc(currentCost,
		comparisonFunc(costSelectingCurrentRing,
			costNotSelectingCurrentRing),
	)
}

func getCost(boss *Character, comparisonFunc func(int, int) int, initialCost int, costWinning bool) int {
	validEquipment := map[string][]Equipment{
		"Weapons": []Equipment{
			Equipment{8, 4, 0},
			Equipment{10, 5, 0},
			Equipment{25, 6, 0},
			Equipment{40, 7, 0},
			Equipment{74, 8, 0},
		},
		"Armor": []Equipment{
			Equipment{13, 0, 1},
			Equipment{31, 0, 2},
			Equipment{53, 0, 3},
			Equipment{75, 0, 4},
			Equipment{102, 0, 5},
		},
		"Rings": []Equipment{
			Equipment{25, 1, 0},
			Equipment{50, 2, 0},
			Equipment{100, 3, 0},
			Equipment{20, 0, 1},
			Equipment{40, 0, 2},
			Equipment{80, 0, 3},
		},
	}

	rings := validEquipment["Rings"]
	bestCost := initialCost
	for _, weapon := range validEquipment["Weapons"] {
		costNoArmor := backtrackingEquipmentCombinations(
			boss,
			[]Equipment{weapon},
			0, 0,
			rings,
			comparisonFunc,
			costWinning,
		)
		bestCost = comparisonFunc(bestCost, costNoArmor)
		for _, armor := range validEquipment["Armor"] {
			costCurrentArmor := backtrackingEquipmentCombinations(
				boss,
				[]Equipment{weapon, armor},
				0, 0,
				rings,
				comparisonFunc,
				costWinning,
			)
			bestCost = comparisonFunc(bestCost, costCurrentArmor)
		}
	}

	return bestCost
}

func part1(boss *Character) {
	minInt := func(a, b int) int {
		return min(a, b)
	}

	minCost := getCost(boss, minInt, math.MaxInt, true)
	fmt.Printf("Part 1: %d\n", minCost)
}

func part2(boss *Character) {
	maxInt := func(a, b int) int {
		if a == math.MaxInt {
			return b
		} else if b == math.MaxInt {
			return a
		}

		return max(a, b)
	}

	maxCost := getCost(boss, maxInt, 0, false)
	fmt.Printf("Part 2: %d\n", maxCost)
}

func parseBoss(input []string) *Character {
	var hitPoints, damage, armor int
	_, err := fmt.Sscanf(input[0], "Hit Points: %d", &hitPoints)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	_, err = fmt.Sscanf(input[1], "Damage: %d", &damage)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	_, err = fmt.Sscanf(input[2], "Armor: %d", &armor)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &Character{
		hitPoints: hitPoints,
		damage:    damage,
		armor:     armor,
	}
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

	bossCharacter := parseBoss(input)
	if bossCharacter == nil {
		fmt.Println("Couldn't find a Boss character")
		os.Exit(1)
	}

	part1(bossCharacter)
	part2(bossCharacter)
}
