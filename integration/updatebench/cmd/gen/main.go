package main

import (
	"log"
	"os"

	"github.com/jalphad/gosqlgen"
)

func main() {
	schema, err := os.ReadFile("../../schema.sql")
	if err != nil {
		log.Fatalf("read schema: %v", err)
	}

	gen := gosqlgen.New().
		WithPackageName("gen").
		WithOutputPath("../../").
		WithPackagePath("github.com/jalphad/gosqlgen/integration/updatebench")

	if err := gen.Parse(string(schema)); err != nil {
		log.Fatalf("parse schema: %v", err)
	}
	if err := gen.GenerateFiles(); err != nil {
		log.Fatalf("generate files: %v", err)
	}
}
