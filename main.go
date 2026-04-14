// Package main is the entry point for the tfplan-summary tool.
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/vetlekise/tfplan-summary-go/parser"
	"github.com/vetlekise/tfplan-summary-go/reader"
	"github.com/vetlekise/tfplan-summary-go/renderer"
)

var planPath string

func init() {
	flag.StringVar(&planPath, "path", "tfplan.json", "Path to your Terraform Plan .json file.")
	flag.Parse()
}

func main() {
	data := planPath

	file, err := reader.ReadPlan(data)
	if err != nil {
		slog.Error("failed to read plan", "err", err)
		os.Exit(1)
	}

	changes, err := parser.ParseChanges(file)
	if err != nil {
		slog.Error("failed to parse plan", "err", err)
		os.Exit(1)
	}

	if len(changes) == 0 {
		fmt.Println("No changes.")
		return
	}

	renderer.RenderTable(changes)
}
