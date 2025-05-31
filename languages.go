package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// LanguageArchitecture represents a language-specific architecture pattern
type LanguageArchitecture struct {
	folders []string
	readme  string
}

// getGoArchitecture returns Go-specific folder structure
func getGoArchitecture() LanguageArchitecture {
	return LanguageArchitecture{
		folders: []string{
			"cmd/api",
			"internal/domain",
			"internal/application",
			"internal/infrastructure",
			"pkg/common",
			"api/http",
			"configs",
			"docs",
			"test",
		},
		readme: "# Go Project Structure\n\n" +
			"- cmd/: Application entrypoints\n" +
			"- internal/: Private application code\n" +
			"- pkg/: Public shared code\n" +
			"- api/: API definitions\n" +
			"- configs/: Configuration files",
	}
}

// getPythonArchitecture returns Python-specific folder structure
func getPythonArchitecture() LanguageArchitecture {
	return LanguageArchitecture{
		folders: []string{
			"src/package_name",
			"src/package_name/core",
			"src/package_name/api",
			"src/package_name/models",
			"tests",
			"docs",
			"configs",
			"scripts",
		},
		readme: "# Python Project Structure\n\n" +
			"- src/: Source code\n" +
			"- tests/: Test files\n" +
			"- docs/: Documentation\n" +
			"- configs/: Configuration files",
	}
}

// createLanguageStructure creates a language-specific architecture
func createLanguageStructure(root string, language string) error {
	var arch LanguageArchitecture

	switch language {
	case "go":
		arch = getGoArchitecture()
	case "python":
		arch = getPythonArchitecture()
	default:
		return fmt.Errorf("unsupported language: %s", language)
	}

	createFolders(root, arch.folders)

	// Create README.md with structure explanation
	readmePath := filepath.Join(root, "README.md")
	return os.WriteFile(readmePath, []byte(arch.readme), 0644)
}
