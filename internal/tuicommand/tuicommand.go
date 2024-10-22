package tuicommand

import (
	"errors"
	"log"

	"github.com/EdgeLordKirito/ChartMaster/internal/tui"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	argLength := len(args)
	if argLength == 0 {
		log.Println("starting full tui")
		return tui.FullTui()
	} else if argLength == 1 {
		filePath := args[0]
		log.Printf("opening partial tui with argument '%s'", filePath)
		return tui.PartialTui(filePath)
	}
	return errors.New("too many args")
}
