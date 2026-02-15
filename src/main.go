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
func aggregateData(file []byte) (rowData []table.Row) {
	// Create a map using the Plan struct
	var result Plan

	// Parse (Unmarshal) the raw data into the map
	json.Unmarshal(file, &result)

	var rows []table.Row
	// Loop through to find your keys and values, then append them to rows
	for _, res := range result.ResourceChanges {
		action := strings.Join(res.Change.Actions, " ")

		// Colorize based on action type
		coloredAction := action
		switch action {
		case "create":
			coloredAction = text.FgGreen.Sprint(action)
		case "delete":
			coloredAction = text.FgRed.Sprint(action)
		case "update":
			coloredAction = text.FgYellow.Sprint(action)
		case "no-op":
			coloredAction = text.FgWhite.Sprint(action)
		case "create delete":
			action = "replace"
			coloredAction = text.FgRed.Sprint(action)
		case "delete create":
			action = "replace"
			coloredAction = text.FgRed.Sprint(action)
		}

		rows = append(rows, table.Row{coloredAction, res.Address})
	}

	return rows
}

// Builds a table using the 'table' package and prints it.
func buildTable(rows []table.Row) {
	tw := table.NewWriter()

	// Configuration and styling
	tw.SortBy([]table.SortBy{{Name: "Action", Mode: table.Dsc}})
	tw.SetStyle(table.StyleLight)

	// Appending
	tw.AppendHeader(table.Row{
		text.Colors{text.FgWhite, text.Bold}.Sprint("Action"),
		text.Colors{text.FgWhite, text.Bold}.Sprint("Addresses")})
	tw.AppendRows(rows)

	fmt.Println(tw.Render())
}
