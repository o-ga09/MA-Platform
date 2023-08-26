package main

import (
	"TDD-practice/ranking"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		os.Exit(3)
	}
	entryfile := args[1]
	playfile := args[2]
	status := ranking.Ranking(entryfile,playfile)

	os.Exit(status)
}