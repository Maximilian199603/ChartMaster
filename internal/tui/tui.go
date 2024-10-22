package tui

import "github.com/EdgeLordKirito/TypeStrengthChart/typechart"

func With(filePath string) error {
	chart, err := typechart.DeserializeFile(filePath)
	if err != nil {
		return err
	}

	// use chart in model
	return nil
}
