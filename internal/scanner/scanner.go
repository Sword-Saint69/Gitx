package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type SecretMatch struct {
	File    string
	Line    int
	Content string
	Type    string
}

var patterns = map[string]*regexp.Regexp{
	"AWS Access Key":     regexp.MustCompile(`AKIA[0-9A-Z]{16}`),
	"AWS Secret Key":     regexp.MustCompile(`(?i)aws(.{0,20})?['"][0-9a-zA-Z\/+]{40}['"]`),
	"GitHub Token":       regexp.MustCompile(`ghp_[0-9a-zA-Z]{36}`),
	"Generic API Key":    regexp.MustCompile(`(?i)api_key(.{0,20})?['"][0-9a-zA-Z]{32,45}['"]`),
	"Slack Webhook":      regexp.MustCompile(`https:\/\/hooks.slack.com\/services\/T[0-9A-Z]{8}\/B[0-9A-Z]{8}\/[0-9a-zA-Z]{24}`),
	"Private Key":        regexp.MustCompile(`-----BEGIN (RSA|EC|DSA|OPENSSH) PRIVATE KEY-----`),
}

func ScanSecrets(root string) ([]SecretMatch, error) {
	var matches []SecretMatch

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

		// Skip binary files
		if isBinary(path) {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return nil // Skip files we can't open
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNum := 0
		for scanner.Scan() {
			lineNum++
			line := scanner.Text()
			for name, re := range patterns {
				if re.MatchString(line) {
					matches = append(matches, SecretMatch{
						File:    path,
						Line:    lineNum,
						Content: strings.TrimSpace(line),
						Type:    name,
					})
				}
			}
		}

		return nil
	})

	return matches, err
}

func isBinary(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	binaryExts := map[string]bool{
		".exe": true, ".dll": true, ".so": true, ".dylib": true,
		".jpg": true, ".png": true, ".gif": true, ".pdf": true,
		".zip": true, ".gz": true, ".tar": true, ".7z": true,
	}
	return binaryExts[ext]
}
