package main

import (
	"errors"
	"flag"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/internal/engine"
	localError "github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"github.com/thiduzz/code-kata-invasion/internal/utils"
	"github.com/thiduzz/code-kata-invasion/tools"
	"log"
	"os"
	"time"
)

func main() {
	var (
		flags       = flag.NewFlagSet("settings", flag.ExitOnError)
		mapFilePath = flags.String(constant.MapFilePathParameter, "", "Path of file that represents the map of cities")
		attackerQty = flags.Uint(constant.AlienQtyParameter, 100, "Total amount of attackers to invade the cities")
		maxMoves    = flags.Uint(constant.MaxMoves, 10000, "Maximum moves necessary")
	)

	err := flags.Parse(os.Args[1:])
	if err != nil {
		throwError(err, flags)
	}

	if err := validateInput(mapFilePath); err != nil {
		throwError(err, flags)
	}
	locations, err := utils.ParseNodes(mapFilePath)
	if err != nil {
		throwError(err, flags)
	}
	randomizer := tools.NewRandomizer(time.Now().UnixMilli())
	engineExecutor := engine.NewEngine(locations, randomizer, *attackerQty, *maxMoves)
	engineExecutor.PrepareAttackers(nodes.NewAttackerFactory(randomizer.RandomName))
	if err := engineExecutor.Start(); err != nil {
		throwError(err, flags)
	}
}

func validateInput(path *string) error {
	if *path == "" {
		return localError.NewCommandError(constant.MapFilePathParameter, "invalid file name")
	}
	return nil
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
