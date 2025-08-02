package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func separateIP(ip string) ([]string, []string) {
	var insideBrackets, outsideBrackets []string
	current := []rune{}

	for _, char := range ip {
		if char == '[' {
			if len(current) == 0 {
				continue
			}
			outsideBrackets = append(outsideBrackets, string(current))
			current = []rune{}
		} else if char == ']' {
			if len(current) == 0 {
				continue
			}
			insideBrackets = append(insideBrackets, string(current))
			current = []rune{}
		}
		current = append(current, char)
	}

	if len(current) != 0 {
		outsideBrackets = append(outsideBrackets, string(current))
	}

	return insideBrackets, outsideBrackets

}

func containsAbba(s string) bool {
	for i := 3; i < len(s); i++ {
		if s[i] == s[i-3] && s[i-1] == s[i-2] && s[i] != s[i-1] {
			return true
		}
	}
	return false
}

func isValidIp(ip string) bool {
	segmentsBetweenBrackets, segmentsOutsideBrackets := separateIP(ip)

	if slices.ContainsFunc(segmentsBetweenBrackets, containsAbba) {
		return false
	}

	return slices.ContainsFunc(segmentsOutsideBrackets, containsAbba)
}

func part1(ipList []string) {
	countSupportTLS := 0

	for _, ip := range ipList {
		if isValidIp(ip) {
			countSupportTLS++
		}
	}

	fmt.Printf("Part 1: %d\n", countSupportTLS)

}

func listAbas(segments []string) []string {
	var output []string

	for _, seg := range segments {
		for i := 2; i < len(seg); i++ {
			if seg[i] == seg[i-2] && seg[i] != seg[i-1] {
				output = append(output, seg[i-2:i+1])
			}
		}
	}

	return output
}

func containsBab(segments []string, aba string) bool {
	expectedBab := string([]byte{aba[1], aba[0], aba[1]})
	for _, seg := range segments {
		if strings.Contains(seg, expectedBab) {
			return true
		}
	}
	return false
}

func part2(ipList []string) {
	countSupportSSL := 0

	for _, ip := range ipList {
		segmentsBetweenBrackets, segmentsOutsideBrackets := separateIP(ip)
		listAba := listAbas(segmentsOutsideBrackets)
		for _, aba := range listAba {
			if containsBab(segmentsBetweenBrackets, aba) {
				countSupportSSL++
				break
			}
		}

	}

	fmt.Printf("Part 2: %d\n", countSupportSSL)
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
	ipList := strings.Split(string(fileContent), "\n")

	part1(ipList)
	part2(ipList)
}
