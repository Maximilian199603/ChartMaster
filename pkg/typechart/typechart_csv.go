package typechart

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

func DeserializeFile(filepath string) (*TypeChart, error) {
	// check if the file exists
	if _, err := os.Stat(filepath); errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	// open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return Deserialize(file)
}

// This function requires that the csv data is of square format
// it does not require that the attacking type set is equal to the defending type set
func Deserialize(r io.Reader) (*TypeChart, error) {
	result := NewTypeChart()

	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	defendingTypes := records[0][1:]

	for i := 1; i < len(records); i++ {
		attackingType := records[i][0]
		for j := 1; j < len(records[i]); j++ {
			effectivenessStr := records[i][j]
			effectiveness, err := strconv.ParseFloat(effectivenessStr, 64)
			if err != nil {
				return nil, err
			}
			defendingType := defendingTypes[j-1]
			result.AddInteraction(attackingType, defendingType, effectiveness)
		}
	}
	return result, nil
}

func SerializeToFile(chart *TypeChart, filepath string) error {
	if _, err := os.Stat(filepath); err == nil {
		return fs.ErrExist
	} else if !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		return err
	}

	csvData, err := Serialize(chart)
	if err != nil {
		return err
	}

	_, err = file.WriteString(csvData)
	return err
}

func Serialize(chart *TypeChart) (string, error) {
	var builder strings.Builder
	writer := csv.NewWriter(&builder)

	// Write the header row
	header := []string{""}
	for defendingType := range chart.chart {
		header = append(header, defendingType)
	}
	if err := writer.Write(header); err != nil {
		return "", err
	}

	// Write each row for attacking types
	for attackingType, defendingMap := range chart.chart {
		row := []string{attackingType}
		for _, defendingType := range header[1:] {
			effectiveness := defendingMap[defendingType]
			row = append(row, fmt.Sprintf("%.1f", effectiveness))
		}
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		return "", err
	}

	return builder.String(), nil
}
