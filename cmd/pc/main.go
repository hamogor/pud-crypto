package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var pcCmd = &cobra.Command{
		Use:   "pc",
		Short: "Pud Crypto CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	pcCmd.AddCommand(versionCmd)
	pcCmd.AddCommand(balancesCmd())
	pcCmd.AddCommand(txCmd())

	err := pcCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
