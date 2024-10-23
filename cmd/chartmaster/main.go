package main

import (
	"os"

	"github.com/EdgeLordKirito/ChartMaster/cmd/chartmaster/appinfo"
	"github.com/EdgeLordKirito/ChartMaster/internal/tuicommand"
	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use: appinfo.AppName,
	}
	//rootCmd.AddCommand(chartCommand())
	rootCmd.AddCommand(tuiCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) // Let Cobra handle printing the error
	}
}

func tuiCommand() *cobra.Command {
	var tuiCmd = &cobra.Command{
		Use:   "tui",
		Short: "Command for opening the TUI",
		Long:  "tui opens the TUI of " + appinfo.AppName,
		Args:  cobra.ExactArgs(1),
		RunE:  tuicommand.Run,
	}

	return tuiCmd
}
