package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/EdgeLordKirito/ChartMaster/cmd/chartmaster/appinfo"
	"github.com/EdgeLordKirito/ChartMaster/internal/appdomain"
	"github.com/EdgeLordKirito/ChartMaster/internal/chartcommand"
	"github.com/EdgeLordKirito/ChartMaster/internal/tuicommand"
	"github.com/spf13/cobra"
)

func main() {

	initLogging()
	var rootCmd = &cobra.Command{
		Use: appinfo.AppName,
	}
	rootCmd.AddCommand(chartCommand())
	rootCmd.AddCommand(tuiCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1) // Let Cobra handle printing the error
	}
}

func initLogging() {
	domainPath, err := appdomain.DomainPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file, err := os.OpenFile(filepath.Join(domainPath, appinfo.AppName+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

}

func chartCommand() *cobra.Command {
	var chartCmd = &cobra.Command{
		Use:   "chart",
		Short: "Command for managing charts",
		Long:  "chart is a CLI tool that allows you to manage the registered charts using various subcommands",
	}

	var addChartCmd = &cobra.Command{
		Use:   "add",
		Short: "add a new chart",
		Long:  "add allows adding a new chart to the list of registered charts",
		Args:  cobra.ExactArgs(2),
		RunE:  chartcommand.AddCommand,
	}

	var updateChartCmd = &cobra.Command{
		Use:   "update",
		Short: "updates a chart",
		Long:  "updates an existing chart with the data of the specified file",
		Args:  cobra.ExactArgs(2),
		RunE:  chartcommand.UpdateCommand,
	}

	var readChartCmd = &cobra.Command{
		Use:   "read",
		Short: "reads the specified charts data",
		Long:  "reads an specific charts adata and outputs it",
		Args:  cobra.ExactArgs(1),
		RunE:  chartcommand.ReadCommand,
	}

	var removeChartCmd = &cobra.Command{
		Use:   "remove",
		Short: "removes the chart",
		Long:  "removes the specific chart from the list of registered charts",
		Args:  cobra.ExactArgs(1),
		RunE:  chartcommand.RemoveCommand,
	}

	var listChartCmd = &cobra.Command{
		Use:   "list",
		Short: "list all registered charts",
		Long:  "list displays all registered charts",
		Args:  cobra.NoArgs,
		RunE:  chartcommand.ListCommand,
	}

	chartCmd.AddCommand(addChartCmd)
	chartCmd.AddCommand(updateChartCmd)
	chartCmd.AddCommand(readChartCmd)
	chartCmd.AddCommand(removeChartCmd)
	chartCmd.AddCommand(listChartCmd)

	return chartCmd
}

func tuiCommand() *cobra.Command {
	var tuiCmd = &cobra.Command{
		Use:   "tui",
		Short: "Command for opening the TUI",
		Long:  "tui opens the TUI of " + appinfo.AppName,
		Args:  cobra.NoArgs,
		Run:   tuicommand.Run,
	}

	var withCmd = &cobra.Command{
		Use:   "with",
		Short: "Open TUi with chart",
		Long:  "Opens the TUI with the specified chart if it is in the registered charts",
		Args:  cobra.ExactArgs(1),
		RunE:  tuicommand.WithCommand,
	}

	tuiCmd.AddCommand(withCmd)

	return tuiCmd
}
