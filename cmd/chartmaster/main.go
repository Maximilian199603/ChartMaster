package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/EdgeLordKirito/ChartMaster/cmd/chartmaster/appinfo"
	"github.com/EdgeLordKirito/ChartMaster/internal/appdomain"
	"github.com/EdgeLordKirito/ChartMaster/internal/tuicommand"
	"github.com/spf13/cobra"
)

func main() {

	initLogging()
	var rootCmd = &cobra.Command{
		Use: appinfo.AppName,
	}
	//rootCmd.AddCommand(chartCommand())
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

func tuiCommand() *cobra.Command {
	var tuiCmd = &cobra.Command{
		Use:   "tui",
		Short: "Command for opening the TUI",
		Long:  "tui opens the TUI of " + appinfo.AppName,
		Args:  cobra.MaximumNArgs(1),
		RunE:  tuicommand.Run,
	}

	return tuiCmd
}
