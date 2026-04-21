package ui

import (
	"fmt"

	"github.com/user/gitx/internal/git"
)

func RenderWeather(w *git.WeatherReport) string {

	icon := "*" // Clear/Standard
	if w.Condition == "Thunderstorm" {
		icon = "!!" 
	} else if w.Condition == "Windy" {
		icon = "~~" 
	} else if w.Condition == "Foggy" {
		icon = "==" 
	}

	content := fmt.Sprintf("%s %s\n%s\n\n", 
		PrimaryStyle.Render(icon), 
		ValueStyle.Render(w.Condition),
		SubtleStyle.Render(w.Summary))

	stats := fmt.Sprintf("%s\n%s\n%s",
		InfoField("TEMP:", fmt.Sprintf("%d°F", w.Temperature)),
		InfoField("WIND:", fmt.Sprintf("%d commits/wk", w.WindSpeed)),
		InfoField("HUMI:", fmt.Sprintf("%d%%", w.Humidity)))

	return Card("Repository Weather", content+stats)
}
