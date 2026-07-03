package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jalphad/gosqlgen"
)

func main() {
	// Generate DTOs from SQL schema
	generateModels()
}

func generateModels() {
	inputPath := "./dump.sql"
	outputPath := "./output"
	packagePath := "github.com/jalphad/gosqlgen/integration/output"
	// Define your SQL schema
	bts, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Could not read file: %s", err.Error())
	}
	schema := string(bts)

	// Create generator
	gen := gosqlgen.New().
		WithPackageName("output").
		WithOutputPath(outputPath).
		WithPackagePath(packagePath)

	// Parse the schema
	if err := gen.Parse(schema); err != nil {
		log.Fatal("Failed to parse schema:", err)
	}

	// Generate Go code
	err = gen.GenerateFiles()
	if err != nil {
		log.Fatal("Failed to generate code:", err)
	}

	fmt.Println("\nModels generated successfully to", outputPath)
}
