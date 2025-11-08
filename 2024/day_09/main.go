package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

func part1(disk string) {
	output := 0
	currentPos := 0
	left, right := 0, len(disk)-1

	if right%2 == 1 {
		right--
	}

	remainingRight := int(disk[right] - '0')

	for left < right {
		if left%2 == 0 {
			blockID := left / 2
			fileSize := int(disk[left] - '0')

			output += blockID * int(int64(fileSize)*int64(currentPos)+
				int64(fileSize)*(int64(fileSize)-1)/2)

			currentPos += fileSize
			left++
		} else {
			freeSpace := int(disk[left] - '0')
			for freeSpace > 0 && left < right {
				if remainingRight == 0 {
					right -= 2
					if right <= left {
						break
					}
					remainingRight = int(disk[right] - '0')
				}

				usedSpaces := min(freeSpace, remainingRight)

				blockID := right / 2
				output += blockID * int(int64(usedSpaces)*int64(currentPos)+
					int64(usedSpaces)*(int64(usedSpaces)-1)/2)

				currentPos += usedSpaces
				freeSpace -= usedSpaces
				remainingRight -= usedSpaces
			}
			left++
		}
	}

	if left == right && remainingRight > 0 {
		blockID := right / 2
		output += blockID * int(int64(remainingRight)*int64(currentPos)+
			int64(remainingRight)*(int64(remainingRight)-1)/2)
		currentPos += remainingRight
	}

	fmt.Printf("Part 1: %d\n", output)
}

type Segment struct {
	blockID, length, start int
}

func part2(disk string) {
	segments := make([]Segment, 0, len(disk))
	nextBlockID, currentPos := 0, 0

	for idx, ch := range disk {
		size := int(ch - '0')
		if idx%2 == 0 {
			segments = append(segments, Segment{blockID: nextBlockID, length: size, start: currentPos})
			nextBlockID++
		} else {
			segments = append(segments, Segment{blockID: -1, length: size, start: currentPos})
		}
		currentPos += size
	}

	for blockID := nextBlockID - 1; blockID >= 0; blockID-- {
		fileIndex := slices.IndexFunc(segments, func(seg Segment) bool {
			return seg.blockID == blockID
		})
		if fileIndex == -1 {
			continue
		}

		file := segments[fileIndex]
		for idx := range fileIndex {
			freeSegment := segments[idx]
			if freeSegment.blockID != -1 || freeSegment.length < file.length {
				continue
			}

			segments[fileIndex].start = freeSegment.start

			if freeSegment.length == file.length {
				segments[idx] = segments[fileIndex]
				segments[fileIndex] = Segment{-1, file.length, file.start}
			} else {
				prevSegment := freeSegment
				segments[idx] = segments[fileIndex]
				segments[fileIndex] = Segment{-1, file.length, file.start}

				newEmptySegment := Segment{-1, prevSegment.length - file.length, prevSegment.start + file.length}
				segments = append(segments[:idx+1], append([]Segment{newEmptySegment}, segments[idx+1:]...)...)
			}

			break
		}
	}

	output := 0
	for _, aSegment := range segments {
		if aSegment.blockID == -1 {
			continue
		}
		sumOfPositions := aSegment.length*aSegment.start + aSegment.length*(aSegment.length-1)/2
		output += aSegment.blockID * sumOfPositions

	}

	fmt.Printf("Part 2: %d\n", output)
}

func main() {
	fileName := flag.String("file", "", "Path to the file")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("file name not provided. Use --file to specify the file path.")
		os.Exit(1)
	}

	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Println("error reading input file:", err)
		os.Exit(1)
	}

	disk := strings.TrimSpace(string(content))
	part1(disk)
	part2(disk)
}
