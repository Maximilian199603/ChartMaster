package tuicommand

import (
	"github.com/EdgeLordKirito/ChartMaster/internal/tui"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	return tui.With(args[0])
}
