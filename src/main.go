package main

import (
	"encoding/json"
	"flag"
	"fmt"
	//"github.com/jedib0t/go-pretty/v6/list"
	//"github.com/jedib0t/go-pretty/v6/progress"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var pathFlag string

// Top level object
type Plan struct {
	ResourceChanges []ResourceChange `json:"resource_changes"`
}

// The items inside the list
type ResourceChange struct {
	Address string `json:"address"`
	Change  Change `json:"change"`
}

// The nested object containing the actions list
type Change struct {
	Actions []string `json:"actions"`
}

func init() {
	flag.StringVar(&pathFlag, "path", "tfplan.json", "Path to your Terraform Plan .json file.")
	flag.Parse()
}

func main() {
	data := pathFlag
	file := validateJson(data)
	rows := aggregateData(file)
	buildTable(rows)
}

// Validates the file extension '.json' of the provided file and reads it.
func validateJson(data string) (file []byte) {

	// Validate extension
	fileExtension := filepath.Ext(data)
	if fileExtension != ".json" {
		panic("File extension must be '.json'!")
	}

	// Open the JSON file
	file, err := os.ReadFile(data)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	return
}

// Aggregates data from the 'Plan' struct, unmarshals the Json file, and appends the data to table rows.
func aggregateData(file []byte) []table.Row {
	var result Plan
	json.Unmarshal(file, &result)

	// Extract raw data
	type Item struct{ Action, Address string }
	var items []Item

	for _, res := range result.ResourceChanges {
		action := strings.Join(res.Change.Actions, " ")

		if action == "no-op" {
			continue
		}

		if action == "create delete" || action == "delete create" {
			action = "replace"
		}
		items = append(items, Item{Action: action, Address: res.Address})
	}

	// Sort alphabetically by raw Action text
	sort.Slice(items, func(i, j int) bool {
		return items[i].Action > items[j].Action
	})

	// Apply colors and build table rows
	var rows []table.Row
	for _, item := range items {
		coloredAction := item.Action
		switch item.Action {
		case "create":
			coloredAction = text.FgGreen.Sprint(item.Action)
		case "delete", "replace":
			coloredAction = text.FgRed.Sprint(item.Action)
		case "update":
			coloredAction = text.FgYellow.Sprint(item.Action)
		case "no-op":
			coloredAction = text.FgWhite.Sprint(item.Action)
		}
		rows = append(rows, table.Row{coloredAction, item.Address})
	}

	return rows
}

// Builds a table using the 'table' package and prints it.
func buildTable(rows []table.Row) {
	tw := table.NewWriter()

	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false},
	})

	// Configuration and styling
	tw.SortBy([]table.SortBy{{Name: "Action", Mode: table.Dsc}})
	tw.SetStyle(table.StyleDefault)
	tw.Style().Options.SeparateRows = true

	// Appending
	tw.AppendHeader(table.Row{
		text.Colors{text.FgWhite}.Sprint("Action"),
		text.Colors{text.FgWhite}.Sprint("Addresses")})
	tw.AppendRows(rows)

	fmt.Println(tw.Render())
}
