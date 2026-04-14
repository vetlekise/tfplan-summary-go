package renderer

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/vetlekise/tfplan-summary-go/parser"
)

// RenderTable renders a table of resource changes to stdout.
func RenderTable(changes []parser.ResourceDiff) {
	tw := table.NewWriter()

	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AutoMerge: false},
	})

	// Configuration and styling
	tw.SortBy([]table.SortBy{{Name: "Action", Mode: table.Dsc}})
	tw.SetStyle(table.StyleDefault)
	tw.Style().Options.SeparateRows = true

	tw.AppendHeader(table.Row{
		text.Colors{text.FgWhite}.Sprint("Action"),
		text.Colors{text.FgWhite}.Sprint("Address")})

	for _, change := range changes {
		coloredAction := change.Action
		switch change.Action {
		case "create":
			coloredAction = text.FgGreen.Sprint(change.Action)
		case "delete", "replace":
			coloredAction = text.FgRed.Sprint(change.Action)
		case "update":
			coloredAction = text.FgYellow.Sprint(change.Action)
		}
		tw.AppendRow(table.Row{coloredAction, change.Address})
	}

	fmt.Println(tw.Render())
}
