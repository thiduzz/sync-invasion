package main

import (
	"errors"
	"flag"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	localError "github.com/thiduzz/code-kata-invasion/internal/errors"
	"log"
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
	if err != nil {
		throwError(err, flags)
	}
}

func throwError(err error, flags *flag.FlagSet) {
	var ce *localError.CommandError
	if errors.As(err, &ce) {
		log.Println(err.Error())
		flags.PrintDefaults()
	} else {
		log.Printf("error when parsing map: %s\n", err.Error())
	}
	os.Exit(1)
}
