package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command line flags
	folderName := flag.String("folder-name", "", "Name of the root folder to create")
	archType := flag.String("type", "", "Architecture type (ddd, clean, layered, mvc)")
	lang := flag.String("lang", "", "Programming language (go, python)")

	// Parse flags
	flag.Parse()

	// Validate required flags
	if *folderName == "" || *archType == "" {
		fmt.Println("Usage: go-arc --folder-name <name> --type <architecture-type> [--lang <language>]")
		fmt.Println("Available architecture types: ddd, clean, layered, mvc")
		fmt.Println("Available languages: go, python")
		os.Exit(1)
	}

	// Create the root folder
	err := os.MkdirAll(*folderName, 0755)
	if err != nil {
		fmt.Printf("Error creating root folder: %v\n", err)
		os.Exit(1)
	}

	// Generate the folder structure based on architecture type
	switch *archType {
	case "ddd":
		createDDDStructure(*folderName)
	case "clean":
		createCleanStructure(*folderName)
	case "layered":
		createLayeredStructure(*folderName)
	case "mvc":
		createMVCStructure(*folderName)
	default:
		fmt.Printf("Unsupported architecture type: %s\n", *archType)
		os.Exit(1)
	}

	// If language is specified, create language-specific structure
	if *lang != "" {
		if err := createLanguageStructure(*folderName, *lang); err != nil {
			fmt.Printf("Error creating language structure: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("Successfully created %s architecture", *archType)
	if *lang != "" {
		fmt.Printf(" with %s language structure", *lang)
	}
	fmt.Printf(" in folder: %s\n", *folderName)
}
