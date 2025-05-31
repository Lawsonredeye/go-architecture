package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// createDDDStructure creates a Domain-Driven Design folder structure
func createDDDStructure(root string) {
	folders := []string{
		"domain/entities",
		"domain/value-objects",
		"domain/aggregates",
		"domain/repositories",
		"domain/services",
		"application/services",
		"application/dtos",
		"infrastructure/persistence",
		"infrastructure/external",
		"interfaces/api",
		"interfaces/web",
		"shared/common",
	}
	createFolders(root, folders)
}

// createCleanStructure creates a Clean Architecture folder structure
func createCleanStructure(root string) {
	folders := []string{
		"entities",
		"usecases",
		"interfaces/controllers",
		"interfaces/presenters",
		"infrastructure/database",
		"infrastructure/external",
		"shared/common",
	}
	createFolders(root, folders)
}

// createLayeredStructure creates a Layered Architecture folder structure
func createLayeredStructure(root string) {
	folders := []string{
		"presentation",
		"business",
		"persistence",
		"database",
		"common",
	}
	createFolders(root, folders)
}

// createMVCStructure creates an MVC Architecture folder structure
func createMVCStructure(root string) {
	folders := []string{
		"models",
		"views",
		"controllers",
		"public",
		"config",
		"utils",
	}
	createFolders(root, folders)
}

// createFolders is a helper function to create multiple folders
func createFolders(root string, folders []string) {
	for _, folder := range folders {
		path := filepath.Join(root, folder)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("Error creating folder %s: %v\n", path, err)
			continue
		}

		// Create a .gitkeep file to preserve empty directories
		gitkeepPath := filepath.Join(path, ".gitkeep")
		if _, err := os.Create(gitkeepPath); err != nil {
			fmt.Printf("Error creating .gitkeep in %s: %v\n", path, err)
		}
	}
}
