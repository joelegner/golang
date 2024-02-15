package main

import (
	"fmt"
	"strings"
)

var index = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

func readFile() string {
	return `
Design metal building foundations | Determine program
    Get input | readFile function
        Modeling Module
            Read input file
            Create model instance
    Process input
        Engineering Module
            Manage design criteria
                Design Criteria Class
                    Return edge_dist
                    Return concrete strength
                    Return rebar yield
                    Return q_a
                    Return T_inc
                    Return T_min
            Obtain model instance
`
}

func calcIndentLevel(line string) int {
	retval := -1
	for strings.HasPrefix(line, "   ") {
		retval++
		line = line[4:]
	}
	return retval
}

func calculateIndex(line string) string {
	lineIndentLevel := calcIndentLevel(line)
	if lineIndentLevel < 0 {
		return "0"
	}
	if lineIndentLevel >= len(index) {
		lineIndentLevel = len(index) - 1
	}
	index[lineIndentLevel]++
	indexList := make([]string, lineIndentLevel+1)
	for i := 0; i <= lineIndentLevel; i++ {
		indexList[i] = fmt.Sprintf("%d", index[i])
	}
	return strings.Join(indexList, "")
}

func splitTextIntoLines(text string) []string {
	return strings.Split(text, "\n")
}

func splitLineIntoFRandDP(text string) []string {
	return strings.Split(text, "|")
}

func main() {
	var designParameter string = ""
	curLevel := 0
	text := readFile()
	for _, line := range splitTextIntoLines(text)[1 : len(strings.Split(text, "\n"))-1] {
		splitLine := splitLineIntoFRandDP(line)
		functionalRequirement := splitLine[0]
		hasDP := (len(splitLine) == 2)
		if hasDP {
			designParameter = splitLine[1]
		}
		lineLevel := calcIndentLevel(functionalRequirement)
		if lineLevel < curLevel {
			// Clear out lower levels when you outdent
			index = append(index[:lineLevel+1], make([]int, len(index)-lineLevel-1)...)
		}
		curLevel = lineLevel
		lineIndex := calculateIndex(functionalRequirement)
		fmt.Printf("FR%s: %s\n", lineIndex, strings.TrimSpace(functionalRequirement))
		if hasDP {
			fmt.Printf("DP%s: %s\n", lineIndex, strings.TrimSpace(designParameter))
		}
	}
}
