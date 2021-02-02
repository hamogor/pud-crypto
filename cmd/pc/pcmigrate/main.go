package main

import (
	"Blockchain/database"
	"fmt"
	"os"
	"time"
)

func main() {
	state, err := database.NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("harry", "harry", 3, ""),
			database.NewTx("harry", "harry", 700, "reward"),
		},
	)

	state.AddBlock(block0)
	block0hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0hash,
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("harry", "babayaga", 2000, ""),
			database.NewTx("harry", "harry", 100, "reward"),
			database.NewTx("babayaga", "harry", 1, ""),
			database.NewTx("babayaga", "caesar", 1000, ""),
			database.NewTx("babayaga", "harry", 50, ""),
			database.NewTx("harry", "harry", 600, "reward"),
		},
	)

	state.AddBlock(block1)
	state.Persist()
}
