package chartvalidation

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/EdgeLordKirito/TypeStrengthChart/typechart"
)

type ChartValidationError struct {
	Internal error
}

func (e *ChartValidationError) Error() string {
	return fmt.Sprintf("chart parse: %v", e.Internal)
}

func ReadAndValidate(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	_, err = typechart.Deserialize(bytes.NewReader(data))
	if err != nil {
		return nil, &ChartValidationError{Internal: err}
	}

	return data, nil
}
