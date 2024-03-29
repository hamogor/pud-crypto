package main

import (
	"fmt"
	"os"
	"Blockchain/database"
	"github.com/spf13/cobra"
	"time"
)

var migrateCmd = func() *cobra.Command {
	var migrateCmd = &cobra.Command {
		Use: "migrate",
		Short: "Migrates the blockchain database according to new business rules.",
		Run: func(cmd *cobra.Command, args []string) {
			dataDir, _ := cmd.Flags().GetString(flagDataDir)

			state, err := database.NewStateFromDisk(dataDir)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			block0 := database.NewBlock(
				database.Hash{},
				0,
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("harry", "harry", 3, ""),
					database.NewTx("harry", "harry", 700, "reward"),
				},
			)

			state.AddBlock(block0)
			block0Hash, err := state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			block1 := database.NewBlock(
				block0Hash,
				1,
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("harry", "ryan", 2000, ""),
					database.NewTx("harry", "harry", 100, "reward"),
					database.NewTx("ryan", "harry", 1, ""),
					database.NewTx("ryan", "charlie", 1000, ""),
					database.NewTx("ryan", "harry", 50, ""),
					database.NewTx("harry", "harry", 600, "reward"),
				},
			)

			state.AddBlock(block1)
			block1Hash, err := state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			block2 := database.NewBlock(
				block1Hash,
				2,
				uint64(time.Now().Unix()),
				[]database.Tx{
					database.NewTx("harry", "harry", 24700, "reward"),
				},
			)

			state.AddBlock(block2)
			_, err = state.Persist()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			} 

		},
	}

	addDefaultRequiredFlags(migrateCmd)

	return migrateCmd
}