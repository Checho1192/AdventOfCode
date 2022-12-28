/*
--- Day 4: Camp Cleanup ---
Space needs to be cleared before the last supplies can be unloaded from the ships, and so several Elves have been assigned the job of cleaning up sections of the camp. Every section has a unique ID number, and each Elf is assigned a range of section IDs.

However, as some of the Elves compare their section assignments with each other, they've noticed that many of the assignments overlap. To try to quickly find overlaps and reduce duplicated effort, the Elves pair up and make a big list of the section assignments for each pair (your puzzle input).

For example, consider the following list of section assignment pairs:

2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
For the first few pairs, this list means:

Within the first pair of Elves, the first Elf was assigned sections 2-4 (sections 2, 3, and 4), while the second Elf was assigned sections 6-8 (sections 6, 7, 8).
The Elves in the second pair were each assigned two sections.
The Elves in the third pair were each assigned three sections: one got sections 5, 6, and 7, while the other also got 7, plus 8 and 9.
This example list uses single-digit section IDs to make it easier to draw; your actual list might contain larger numbers. Visually, these pairs of section assignments look like this:

.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8
Some of the pairs have noticed that one of their assignments fully contains the other. For example, 2-8 fully contains 3-7, and 6-6 is fully contained by 4-6. In pairs where one assignment fully contains the other, one Elf in the pair would be exclusively cleaning sections their partner will already be cleaning, so these seem like the most in need of reconsideration. In this example, there are 2 such pairs.

In how many assignment pairs does one range fully contain the other?
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readLines("day4.txt")
	parsedLines := parsedInput(lines)
	fmt.Println("Part 1:", part1(parsedLines))
	fmt.Println("Part 2:", part2(parsedLines))

}

func part1(input [][][]int) (count int) {
	for _, pairs := range input {
		if doesPairContainOther(pairs[0], pairs[1]) || doesPairContainOther(pairs[1], pairs[0]) {
			count++
		}
	}
	return count
}

func part2(input [][][]int) (count int) {
	for _, pairs := range input {
		if doesOverlap(pairs[0], pairs[1]) {
			count++
		}
	}
	return count
}

func doesPairContainOther(pair1, pair2 []int) bool {
	return pair1[0] >= pair2[0] && pair1[1] <= pair2[1]
}

func doesOverlap(pair1, pair2 []int) bool {
	if pair1[0] > pair2[0] {
		pair1, pair2 = pair2, pair1
	}
	return pair1[1] >= pair2[0]

}
func readLines(path string) (lines []string) {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	return lines
}

func parsedInput(lines []string) (parsedLines [][][]int) {
	for _, value := range lines {
		separated := strings.Split(value, ",")
		pair1 := strings.Split(separated[0], "-")
		pair2 := strings.Split(separated[1], "-")

		intPair1 := make([]int, 2)
		intPair2 := make([]int, 2)
		for i := 0; i < 2; i++ {
			intPair1[i], _ = strconv.Atoi(pair1[i])
			intPair2[i], _ = strconv.Atoi(pair2[i])
		}
		pairs := [][]int{intPair1, intPair2}
		parsedLines = append(parsedLines, pairs)
	}
	return parsedLines
}
