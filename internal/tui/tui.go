package tui

import (
	"github.com/EdgeLordKirito/ChartMaster/pkg/typechart"
	"github.com/charmbracelet/huh"
	"log"
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
	}

	return nil
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

	// use selection to generate the table of effectivenesses

	// Add table

	return *form
}

func typeOptions(chart typechart.TypeChart) []huh.Option[string] {
	types := chart.DefendingTypes()
	options := make([]huh.Option[string], 0, len(types))
	for typeName := range types {
		options = append(options, huh.NewOption(typeName, typeName))
	}
	return options
}
