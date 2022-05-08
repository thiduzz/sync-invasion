package utils

import (
	"bufio"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"os"
	"strings"
)

//ParseNodes Utility function - For reading the map file and convert it into a collection of locations
// Assumption: even for the links specified within the file are considered, meaning, if a city A points to city B,
// i won't be creating a link from city B to city A, only a link from city A to city B
func ParseNodes(filePath *string) (*nodes.LocationCollection, error) {
	if *filePath == "" {
		return nil, errors.NewCommandError(constant.MapFilePathParameter, "missing parameter")
	}

	file, err := os.Open(*filePath)
	if err != nil {
		return nil, err
	}
	//guarantees the file will be closed when function is finished/application crashes
	defer file.Close()
	locations := nodes.NewLocationCollection()
	scanner := bufio.NewScanner(file)
	id := uint(1)
	for scanner.Scan() {

		pieces := strings.Split(scanner.Text(), " ")

		if len(pieces) == 0 || (len(pieces) == 1 && pieces[0] == "") {
			return nil, errors.NewUtilsError("ParseMap", "invalid row in parsed file")
		}

		locationName := pieces[0]
		location := nodes.NewLocation(id, locationName)
		locations.Add(location)
		id++
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.NewUtilsErrorWrap("ParseMap", err)
	}

	return locations, nil
}
