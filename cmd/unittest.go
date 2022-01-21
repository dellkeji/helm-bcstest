package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/dellkeji/helm-bcstest/pkg/printer"
)

var unittestCmd = &cobra.Command{
	Use:   "unittest",
	Short: "unittest for helm3 charts",
	Run:   unittestRun,
}

func unittestRun(cmd *cobra.Command, args []string) {
	printer := printer.NewPrinter(os.Stdout)
	printer.Print(printer.HighlightPrint("start test ..."), 0)
}

func init() {
	// unittestCmd.PersistentFlags()
	rootCmd.AddCommand(unittestCmd)
}
