// Package main is the entry point for the tfplan-summary tool.
package main

import (
	"flag"

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
	file := reader.ReadPlan(data)
	rows := parser.ParseChanges(file)
	renderer.RenderTable(rows)
}
