package git

import (
	"os"
	"path/filepath"
	"strings"
)

type LanguageStats struct {
	Files int
	Lines int
}

func GetLOC(root string) (map[string]*LanguageStats, error) {
	stats := make(map[string]*LanguageStats)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == "node_modules" || info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}

		// Detect language by extension
		ext := strings.ToLower(filepath.Ext(path))
		if ext == "" {
			return nil
		}

		lang := "Unknown"
		switch ext {
		case ".go":
			lang = "Go"
		case ".js", ".jsx":
			lang = "JavaScript"
		case ".ts", ".tsx":
			lang = "TypeScript"
		case ".py":
			lang = "Python"
		case ".rs":
			lang = "Rust"
		case ".md":
			lang = "Markdown"
		case ".html":
			lang = "HTML"
		case ".css":
			lang = "CSS"
		case ".json":
			lang = "JSON"
		case ".toml":
			lang = "TOML"
		case ".yaml", ".yml":
			lang = "YAML"
		default:
			return nil // Skip unknown for now to keep it clean
		}

		if stats[lang] == nil {
			stats[lang] = &LanguageStats{}
		}

		lines, err := countLines(path)
		if err != nil {
			return nil
		}

		stats[lang].Files++
		stats[lang].Lines += lines

		return nil
	})

	return stats, err
}

