package tuicommand

import (
	"fmt"

	"github.com/spf13/cobra"
)

type WrappedError struct {
	internal error
}

func Wrap(input error) *WrappedError {
	return &WrappedError{internal: input}
}

func (e *WrappedError) Error() string {
	return fmt.Sprintf("Wrapping error: %v", e.internal)
}

func (e *WrappedError) UnWrap() error {
	return e.internal
}

func WithCommand(cmd *cobra.Command, args []string) error {
	chartName := args[0]
	fmt.Printf("Opening TUI with chart: %s", chartName)
	return nil
}

func Run(cmd *cobra.Command, args []string) {

}
