package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	salt                 = "abc"
	stretch              = 2016
	requiredKeys         = 64
	repeatedTripletLimit = 1000
)

func md5Hex(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func calculateStepHash(step int) string {
	str := salt + strconv.Itoa(step)
	return md5Hex(str)
}
func calculateStretchedHash(step int) string {
	hash := calculateStepHash(step)
	for range stretch {
		hash = md5Hex(hash)
	}
	return hash
}

func getTriplet(s string) rune {
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}

		if count == 3 {
			return rune(s[i])
		}
	}
	return 0
}

func findIndexesValidKeys(getHash func(int) string) []int {
	steps := 0
	stepsValidIndexes := []int{}
	foundTriplets := make(map[int]rune)

	for ; len(stepsValidIndexes) < requiredKeys; steps++ {
		hash := getHash(steps)

		for step, val := range foundTriplets {
			if steps-step > repeatedTripletLimit {
				delete(foundTriplets, step)
			} else if strings.Contains(hash, strings.Repeat(string(val), 5)) {
				stepsValidIndexes = append(stepsValidIndexes, step)
				delete(foundTriplets, step)
			}
		}

		if firstTriplet := getTriplet(hash); firstTriplet != 0 {
			foundTriplets[steps] = firstTriplet
		}
	}
	return stepsValidIndexes
}

func part1() {
	stepsValidIndexes := findIndexesValidKeys(calculateStepHash)
	sort.Ints(stepsValidIndexes)
	fmt.Printf("Part 1: %d\n", stepsValidIndexes[requiredKeys-1])
}

func part2() {
	stepsValidIndexes := findIndexesValidKeys(calculateStretchedHash)
	sort.Ints(stepsValidIndexes)
	fmt.Printf("Part 2: %d\n", stepsValidIndexes[requiredKeys-1])
}

func main() {
	part1()
	part2()
}
