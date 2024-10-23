package tui

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/EdgeLordKirito/ChartMaster/pkg/typechart"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

var (
	typeSelection []string
)

func With(filePath string) error {
	log.Printf("validating file '%s'", filePath)
	chart, err := typechart.DeserializeFile(filePath)
	if err != nil {
		log.Printf("encountered error: '%v'", err)
		return err
	}

	form := setupForm(*chart)
	err = form.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}

	weaknessTable := typechart.NewWeaknesstable(*chart, typeSelection...)

	presentabletable := buildTable(*weaknessTable)
	p := tea.NewProgram(presentabletable)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
	// err is nil
	return err
}

func setupForm(chart typechart.TypeChart) huh.Form {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Types").
				Value(&typeSelection).
				Options(typeOptions(chart)...),
		),
	)

	return *form
}

func typeOptions(chart typechart.TypeChart) []huh.Option[string] {
	types := chart.DefendingTypes()
	var list []string
	for k := range types {
		list = append(list, k)
	}
	sort.Strings(list)

	options := make([]huh.Option[string], 0, len(types))
	for _, typeName := range list {
		options = append(options, huh.NewOption(typeName, typeName))
	}
	return options
}

func buildTable(t typechart.WeaknessTable) tea.Model {
	cols := generateColumns(t)
	rows := generateRows(t)
	height := len(rows)
	table := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(height+1),
	)

	return model{table: table}
}

func generateColumns(input typechart.WeaknessTable) []table.Column {
	converted := input.AsMap()
	var cols []table.Column
	cols = append(cols, table.Column{Title: "index", Width: 8})
	//sort cols
	var titles []float64
	for k := range converted {
		titles = append(titles, k)
	}
	sort.Float64s(titles)
	for _, v := range titles {
		width := longest(converted[v])
		if width < 10 {
			width = 10
		}
		cols = append(cols, table.Column{Title: presentableFloat(v), Width: width})
	}
	return cols
}

func generateRows(weaknesses typechart.WeaknessTable) []table.Row {
	converted := weaknesses.AsMap()
	var titles []float64
	for k := range converted {
		titles = append(titles, k)
	}
	sort.Float64s(titles)

	colNumber := len(titles) + 1
	var rows []table.Row

	// Create a map to hold the maximum length of each weakness list
	maxLen := 0
	for _, title := range titles {
		if len(converted[title]) > maxLen {
			maxLen = len(converted[title])
		}
	}

	// Populate rows based on sorted keys
	for i := 0; i < maxLen; i++ {
		// Create a new row with empty strings
		row := make([]string, colNumber)
		row[0] = fmt.Sprintf("%d", i+1) // Set the first column to the index (1-based)

		// Fill the row with the corresponding weaknesses
		for j, title := range titles {
			if i < len(converted[title]) {
				row[j+1] = converted[title][i] // Fill with the appropriate weakness
			} else {
				row[j+1] = "" // Fill with empty string if no more weaknesses
			}
		}

		// Append the populated row to the rows slice
		rows = append(rows, row)
	}

	return rows
}

func longest(input []string) int {
	max := -1
	for _, v := range input {
		length := len([]rune(v))
		if length > max {
			max = length
		}
	}
	return max
}

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "crtl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.table.View() + "\n press 'q' to quit.\n"
}

func presentableFloat(f float64) string {
	if math.IsNaN(f) || math.IsInf(f, 1) || math.IsInf(f, -1) {
		return fmt.Sprintf("%v", f) // return as is for NaN and Infinity
	}

	if f == 0 {
		return "0" + "\u00D7" // return "0" explicitly for zero
	}

	// Check for significant decimal places
	// First, format to a max of 5 decimal places
	str := fmt.Sprintf("%.5f", f)

	// Remove trailing zeros
	str = strings.TrimRight(str, "0")
	str = strings.TrimRight(str, ".")

	// Check if the result fits in the required length
	if len(str) > 8 {
		str = fmt.Sprintf("%.2e", f) // ensure it uses scientific notation
	}

	// add multiply symbol
	return str + "\u00D7"
}
