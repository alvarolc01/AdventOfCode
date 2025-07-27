package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	AvailableSpace = 70000000
	RequiredSpace  = 30000000
)

type Directory struct {
	name     string
	size     int
	children []*Directory
	isFile   bool
}

func (d *Directory) GetSize() int {
	if d.isFile {
		return d.size
	}

	if d.size != 0 {
		return d.size
	}

	size := 0
	for _, child := range d.children {
		size += child.GetSize()
	}
	d.size = size
	return size
}

func (d *Directory) AddChild(child *Directory) {
	if d.children == nil {
		d.children = make([]*Directory, 0)
	}
	d.children = append(d.children, child)
}

func (d *Directory) GetChildrenDirectoriesSizes() []int {
	sizes := []int{}
	for _, child := range d.children {
		if child.isFile == false {
			childDirsSizes := child.GetChildrenDirectoriesSizes()
			sizes = append(sizes, childDirsSizes...)
		}
	}
	sizes = append(sizes, d.GetSize())
	return sizes
}

func generateFileSystem(commands []string) *Directory {
	createdDirectories := make(map[string]*Directory)
	curr := []string{}
	currDir := &Directory{
		name:   "/",
		isFile: false,
	}

	createdDirectories["/"] = currDir

	for _, command := range commands {
		if strings.HasPrefix(command, "$ cd") {
			if strings.HasSuffix(command, "/") {
				curr = []string{}
				currDir = createdDirectories["/"]
			} else if strings.HasSuffix(command, "..") {
				curr = curr[:len(curr)-1]
				currDir = createdDirectories[strings.Join(curr, "/")]
			} else {
				fields := strings.Fields(command)
				curr = append(curr, fields[2])
				currDir = createdDirectories[strings.Join(curr, "/")]
			}

		} else if !strings.HasPrefix(command, "$ ls") {
			fields := strings.Fields(command)
			size := fields[0]
			name := fields[1]
			newFile := &Directory{
				name: name,
			}

			if size != "dir" {
				fileSize, _ := strconv.Atoi(size)
				newFile.size = fileSize
				newFile.isFile = true
			} else {
				curr = append(curr, name)
				createdDirectories[strings.Join(curr, "/")] = newFile
				curr = curr[:len(curr)-1]
			}
			currDir.AddChild(newFile)
		}
	}
	return createdDirectories["/"]
}

func part1(root *Directory) {
	dirSizes := root.GetChildrenDirectoriesSizes()

	totalSum := 0
	for _, n := range dirSizes {
		if n < 1e5 {
			totalSum += n
		}
	}

	fmt.Printf("Part 1: %d\n", totalSum)
}

func part2(root *Directory) {
	dirSizes := root.GetChildrenDirectoriesSizes()

	sort.Ints(dirSizes)

	remainingSpace := AvailableSpace - root.size
	freedSpaceRequired := RequiredSpace - remainingSpace

	freedSpace := 0
	for _, n := range dirSizes {
		if n >= freedSpaceRequired {
			freedSpace = n
			break
		}
	}
	fmt.Printf("Part 2: %d\n", freedSpace)
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
	input := strings.Split(string(fileContent), "\n")
	rootDir := generateFileSystem(input)

	part1(rootDir)
	part2(rootDir)
}
