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

// getJavaArchitecture returns Java-specific folder structure
func getJavaArchitecture() LanguageArchitecture {
	return LanguageArchitecture{
		folders: []string{
			"src/main/java",
			"src/main/resources",
			"src/test/java",
			"src/test/resources",
			"target",
			"docs",
			"config",
		},
		readme: "# Java Project Structure (Maven Standard)\n\n" +
			"- src/main/java/: Application source code\n" +
			"- src/main/resources/: Configuration files and resources\n" +
			"- src/test/java/: Test source code\n" +
			"- src/test/resources/: Test configuration and resources\n" +
			"- target/: Compiled output\n" +
			"- docs/: Documentation\n" +
			"- config/: Additional configuration files\n",
	}
}

// getJavaScriptArchitecture returns JavaScript-specific folder structure
func getJavaScriptArchitecture() LanguageArchitecture {
	return LanguageArchitecture{
		folders: []string{
			"src",
			"src/components",
			"src/services",
			"src/utils",
			"src/config",
			"src/types",
			"src/middleware",
			"src/__tests__",
			"public",
		},
		readme: "# JavaScript Project Structure\n\n" +
			"- src/: Source code\n" +
			"  - components/: UI components\n" +
			"  - services/: Business logic and API calls\n" +
			"  - utils/: Utility functions\n" +
			"  - config/: Configuration files\n" +
			"  - types/: Type definitions\n" +
			"  - middleware/: Middleware functions\n" +
			"  - __tests__/: Test files\n" +
			"- public/: Static files\n",
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
	case "java":
		arch = getJavaArchitecture()
	case "javascript":
		arch = getJavaScriptArchitecture()
	default:
		return fmt.Errorf("unsupported language: %s. Supported languages: go, python, java, javascript", language)
	}

	createFolders(root, arch.folders)

	// Create README.md with structure explanation
	readmePath := filepath.Join(root, "README.md")
	return os.WriteFile(readmePath, []byte(arch.readme), 0644)
}
