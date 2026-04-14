// Package parser parses Terraform plan JSON into structured resource change data.
package parser

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// Plan represents the top-level structure of a Terraform plan JSON file.
type Plan struct {
	ResourceChanges []ResourceChange `json:"resource_changes"`
}

// ResourceChange represents a single resource entry in the plan.
type ResourceChange struct {
	Address string `json:"address"`
	Change  Change `json:"change"`
}

// Change holds the list of actions to be applied to a resource.
type Change struct {
	Actions []string `json:"actions"`
}

// ResourceDiff represents a parsed resource change with its action and address.
type ResourceDiff struct {
	Action  string
	Address string
}

// ParseChanges parses a Terraform plan JSON and returns a list of resource changes.
func ParseChanges(file []byte) ([]ResourceDiff, error) {
	var result Plan

	if err := json.Unmarshal(file, &result); err != nil {
		return nil, fmt.Errorf("parse plan: %w", err)
	}

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

	return items, nil
}
