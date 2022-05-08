package utils

import (
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"os"
)

func ParseNodes(filePath *string) (*nodes.LocationCollection, error) {
	if *filePath == "" {
		return nil, errors.NewCommandError(constant.MapFilePathParameter, "missing parameter")
	}
	_, err := os.Open(*filePath)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
