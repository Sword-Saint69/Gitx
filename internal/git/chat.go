package git

import (
	"fmt"
	"strings"
)

func AnswerQuestion(path, question string) (string, error) {
	q := strings.ToLower(question)

	if strings.Contains(q, "contributor") || strings.Contains(q, "author") || strings.Contains(q, "who") {
		stats, _ := GetStats(path)
		return fmt.Sprintf("There are %d contributors in this repository. The main author has committed most recently at %s.", stats.Contributors, stats.LastCommit.Format("Jan 2006")), nil
	}

	if strings.Contains(q, "lines") || strings.Contains(q, "loc") || strings.Contains(q, "code") {
		loc, _ := GetLOC(path)
		summary := []string{}
		for lang, stat := range loc {
			summary = append(summary, fmt.Sprintf("%s (%d lines)", lang, stat.Lines))
		}
		return "This project has a diverse stack: " + strings.Join(summary, ", ") + ".", nil
	}

	if strings.Contains(q, "change") || strings.Contains(q, "yesterday") || strings.Contains(q, "recent") {
		standup, _ := GetStandup(path)
		if len(standup) > 0 && len(standup[0].Commits) > 0 {
			return fmt.Sprintf("In the last 24 hours, you've worked on: %s", strings.Join(standup[0].Commits, "; ")), nil
		}
		return "No major changes in the last 24 hours. A quiet day!", nil
	}

	if strings.Contains(q, "risk") || strings.Contains(q, "hotspot") || strings.Contains(q, "mess") {
		hotspots, _ := GetHotspots(path)
		if len(hotspots) > 0 {
			return fmt.Sprintf("The most complex areas are %s. You might want to review them for technical debt.", hotspots[0].File), nil
		}
	}

	return "I'm not sure about that specifically, but I can tell you about stats, LOC, recent changes, or code risks!", nil
}
