package parser

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

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

// Aggregates data from the 'Plan' struct, unmarshals the Json file, and appends the data to table rows.
func ParseChanges(file []byte) []table.Row {
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
