package main

import (
	"Blockchain/node"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Launches the PC node and its HTTP API.",
		Run: func(cmd *cobra.Command, args []string) {
			dataDir, _ := cmd.Flags().GetString(flagDataDir)
			port, _ := cmd.Flags().GetUint64(flagPort)

			fmt.Println("Launching PC node and its HTTP API...")

			
			bootstrap := node.NewPeerNode(
				"172.31.21.72",
				8080,
				true,
				true,
			)

			n := node.New(dataDir, port, bootstrap)
			err := n.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	addDefaultRequiredFlags(runCmd)
	runCmd.Flags().Uint64(flagPort, node.DefaultHTTPPort, "exposed HTTP port for communication with peers")

	return runCmd
}
