package utils

import (
	"bufio"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"os"
	"strings"
)

func ParseNodes(filePath *string) (*nodes.LocationCollection, error) {
	if *filePath == "" {
		return nil, errors.NewCommandError(constant.MapFilePathParameter, "missing parameter")
	}
	file, err := os.Open(*filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	cities := nodes.NewLocationCollection()
	scanner := bufio.NewScanner(file)
	id := uint(1)
	for scanner.Scan() {

		pieces := strings.Split(scanner.Text(), " ")

		if len(pieces) == 0 || (len(pieces) == 1 && pieces[0] == "") {
			return nil, errors.NewUtilsError("ParseMap", "invalid row in parsed file")
		}

		cityName := pieces[0]
		cities.Add(nodes.NewLocation(id, cityName))
		id++
	}

	return cities, nil
}
