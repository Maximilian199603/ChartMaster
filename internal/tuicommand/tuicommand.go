package tuicommand

import (
	"log"

	"github.com/EdgeLordKirito/ChartMaster/internal/tui"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	log.Printf("opening tui with arg '%s'", args[0])
	return tui.With(args[0])
}
