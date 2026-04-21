package git

import (
	"fmt"
	"os"
	"strings"
)

var ignoreTemplates = map[string]string{
	"go": `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
bin/
# Test binary, built with 'go test -c'
*.test
# Output of the go coverage tool, specifically 'go test -coverprofile'
*.out
go.sum`,
	"node": `node_modules/
npm-debug.log
yarn-error.log
.env
dist/
build/`,
	"python": `__pycache__/
*.py[cod]
*$py.class
.env
venv/
ENV/
build/
dist/`,
	"rust": `target/
**/*.rs.bk
Cargo.lock`,
}

func GenerateIgnore(lang string, appendMode bool) error {
	content, ok := ignoreTemplates[strings.ToLower(lang)]
	if !ok {
		return fmt.Errorf("no template found for language: %s", lang)
	}

	mode := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if appendMode {
		mode = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		content = "\n" + content
	}

	f, err := os.OpenFile(".gitignore", mode, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}
