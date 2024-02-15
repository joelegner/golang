package main

import (
	"fmt"
	"strings"
)

var index = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

var text = `
Determine
    Get input
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
		return "FR0"
	}
	if lineIndentLevel >= len(index) {
		lineIndentLevel = len(index) - 1
	}
	index[lineIndentLevel]++
	indexList := make([]string, lineIndentLevel+1)
	for i := 0; i <= lineIndentLevel; i++ {
		indexList[i] = fmt.Sprintf("%d", index[i])
	}
	return "FR" + strings.Join(indexList, "")
}

func main() {
	curLevel := 0
	for _, line := range strings.Split(text, "\n")[1 : len(strings.Split(text, "\n"))-1] {
		lineLevel := calcIndentLevel(line)
		if lineLevel < curLevel {
			// Clear out lower levels when you outdent
			index = append(index[:lineLevel+1], make([]int, len(index)-lineLevel-1)...)
		}
		curLevel = lineLevel
		lineIndex := calculateIndex(line)
		fmt.Printf("%s: %s\n", lineIndex, strings.TrimSpace(line))
	}
}
