package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	PlayerStartingMana int = 500
	PlayerStartingHP   int = 50
	MaxTurns           int = 20
	ShieldArmor        int = 7
)

type Character struct {
	hitPoints int
	damage    int
	mana      int
}

func (c *Character) DeepCopy() *Character {
	return &Character{
		hitPoints: c.hitPoints,
		damage:    c.damage,
		mana:      c.mana,
	}
}

type Effect struct {
	duration        int
	damage          int
	armor           int
	heal            int
	regeneratedMana int
}

type Attack struct {
	name   string
	cost   int
	effect Effect
}

var validAttacks = []Attack{
	{name: "Magic Missile", cost: 53, effect: Effect{duration: 1, damage: 4}},
	{name: "Drain", cost: 73, effect: Effect{duration: 1, damage: 2, heal: 2}},
	{name: "Shield", cost: 113, effect: Effect{duration: 6, armor: 7}},
	{name: "Poison", cost: 173, effect: Effect{duration: 6, damage: 3}},
	{name: "Recharge", cost: 229, effect: Effect{duration: 5, regeneratedMana: 101}},
}

func applyEffect(effect Effect, boss, player *Character) {
	boss.hitPoints -= effect.damage
	player.hitPoints += effect.heal
	player.mana += effect.regeneratedMana
}

func backtrackingMinManaCost(boss, player *Character, spentMana, turn int, currentEffects map[string]Effect, secondPart bool) int {
	if turn >= MaxTurns {
		return math.MaxInt
	}

	if secondPart && turn%2 == 0 {
		player.hitPoints--
		if player.hitPoints == 0 {
			return math.MaxInt
		}
	}

	if player.mana <= 0 {
		return math.MaxInt
	}

	_, useShield := currentEffects["Shield"]
	for key, val := range currentEffects {
		applyEffect(val, boss, player)
		val.duration--
		if val.duration == 0 {
			delete(currentEffects, key)
		} else {
			currentEffects[key] = val
		}
	}

	player.hitPoints = min(player.hitPoints, 50)

	if turn%2 == 1 {
		playerReceivedDamage := boss.damage
		if useShield {
			playerReceivedDamage -= ShieldArmor
		}
		player.hitPoints -= max(1, playerReceivedDamage)
	}

	if boss.hitPoints <= 0 {
		return spentMana
	} else if player.mana <= 0 || player.hitPoints <= 0 {
		return math.MaxInt
	}

	if turn%2 != 0 {
		return backtrackingMinManaCost(boss.DeepCopy(), player.DeepCopy(), spentMana, turn+1, currentEffects, secondPart)
	}

	cost := math.MaxInt
	for _, attack := range validAttacks {
		if _, currentlyApplied := currentEffects[attack.name]; !currentlyApplied {
			attackEffectsCopy := make(map[string]Effect)
			for k, v := range currentEffects {
				attackEffectsCopy[k] = v
			}

			attackEffectsCopy[attack.name] = attack.effect
			player.mana -= attack.cost
			currCost := backtrackingMinManaCost(boss.DeepCopy(), player.DeepCopy(), spentMana+attack.cost, turn+1, attackEffectsCopy, secondPart)
			cost = min(cost, currCost)
			player.mana += attack.cost
		}

	}
	return cost
}

func part1(boss *Character) {

	player := &Character{
		hitPoints: PlayerStartingHP,
		mana:      PlayerStartingMana,
	}

	minCost := backtrackingMinManaCost(boss, player, 0, 0, map[string]Effect{}, false)
	fmt.Printf("Part 1: %d\n", minCost)
}

func part2(boss *Character) {

	player := &Character{
		hitPoints: PlayerStartingHP,
		mana:      PlayerStartingMana,
	}

	minCost := backtrackingMinManaCost(boss, player, 0, 0, map[string]Effect{}, true)
	fmt.Printf("Part 2: %d\n", minCost)
}

func parseBoss(input []string) (*Character, error) {
	var hitPoints, damage int
	_, err := fmt.Sscanf(input[0], "Hit Points: %d", &hitPoints)
	if err != nil {
		return nil, err
	}
	_, err = fmt.Sscanf(input[1], "Damage: %d", &damage)
	if err != nil {
		return nil, err
	}

	return &Character{
		hitPoints: hitPoints,
		damage:    damage,
	}, nil
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

	bossCharacter, err := parseBoss(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	part1(bossCharacter)
	part2(bossCharacter)
}
