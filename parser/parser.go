package parser

import (
	"encoding/json"
	"sort"
	"strings"
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

// ResourceDiff represents a parsed resource change with its action and address.
type ResourceDiff struct {
	Action  string
	Address string
}

// ParseChanges parses a Terraform plan JSON and returns a list of resource changes.
func ParseChanges(file []byte) []ResourceDiff {
	var result Plan
	json.Unmarshal(file, &result)

	var items []ResourceDiff

	for _, res := range result.ResourceChanges {
		action := strings.Join(res.Change.Actions, " ")

		if action == "no-op" {
			continue
		}

		if action == "create delete" || action == "delete create" {
			action = "replace"
		}
		items = append(items, ResourceDiff{Action: action, Address: res.Address})
	}

	// Sort alphabetically by raw Action text
	sort.Slice(items, func(i, j int) bool {
		return items[i].Action > items[j].Action
	})

	return items
}
