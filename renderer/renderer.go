package renderer

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	//"github.com/jedib0t/go-pretty/v6/list"
	//"github.com/jedib0t/go-pretty/v6/progress"
)

// Builds a table using the 'table' package and prints it.
func RenderTable(rows []table.Row) {
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
		text.Colors{text.FgWhite}.Sprint("Address")})
	tw.AppendRows(rows)

	fmt.Println(tw.Render())
}
