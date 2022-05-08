package main

import (
	"code-kata-invasion/internal/constant"
	"flag"
	"os"
)

func main() {
	var (
		flags       = flag.NewFlagSet("settings", flag.ExitOnError)
		mapFilePath = flags.String(constant.MapFilePathParameter, "", "Path of file that represents the map of cities")
		alientQty   = flags.Uint(constant.AlienQtyParameter, 100, "Total amount of aliens to invade")
		maxMoves    = flags.Uint(constant.MaxMoves, 10000, "Maximum moves necessary")
	)

	err := flags.Parse(os.Args[1:])
}
