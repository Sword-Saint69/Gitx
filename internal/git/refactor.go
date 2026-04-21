package git

import (
	"bufio"
	"os"
	"strings"
)

type RefactorSuggestion struct {
	File     string
	Context  string // Function name
	Reason   string
	Priority string // LOW, HIGH
}

func AnalyzeRefactor(filePath string) ([]RefactorSuggestion, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var suggestions []RefactorSuggestion
	scanner := bufio.NewScanner(file)
	
	lineCount := 0
	inFunc := false
	funcLineStart := 0
	currentFunc := ""
	
	for scanner.Scan() {
		lineCount++
		line := strings.TrimSpace(scanner.Text())
		
		if strings.HasPrefix(line, "func ") {
			inFunc = true
			funcLineStart = lineCount
			currentFunc = strings.TrimPrefix(line, "func ")
			if idx := strings.Index(currentFunc, "("); idx != -1 {
				currentFunc = currentFunc[:idx]
			}
		}
		
		if inFunc && line == "}" {
			length := lineCount - funcLineStart
			if length > 40 {
				suggestions = append(suggestions, RefactorSuggestion{
					File:     filePath,
					Context:  currentFunc,
					Reason:   "Function is too long ( > 40 lines). Consider extracting sub-functions.",
					Priority: "HIGH",
				})
			}
			inFunc = false
		}

		// Rule 2: Deep Nesting (Heuristic)
		if strings.Contains(line, "{") && (strings.Contains(line, "if") || strings.Contains(line, "for")) {
			indent := len(scanner.Text()) - len(strings.TrimLeft(scanner.Text(), " \t"))
			if indent > 12 {
				suggestions = append(suggestions, RefactorSuggestion{
					File:     filePath,
					Context:  currentFunc,
					Reason:   "Deep indentation detected. Use early returns to flatten the logic.",
					Priority: "LOW",
				})
			}
		}
	}

	return suggestions, nil
}
