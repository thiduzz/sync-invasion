package utils

import (
	"bufio"
	"fmt"
	"github.com/thiduzz/code-kata-invasion/internal/constant"
	"github.com/thiduzz/code-kata-invasion/internal/errors"
	"github.com/thiduzz/code-kata-invasion/internal/nodes"
	"os"
	"strings"
)

//ParseNodes Utility function - For reading the map file and convert it into a collection of locations
// Assumption: I assume that in a specific direction there might be more than one location and
// that location needs to keep a reference of all the inbound locations to it
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

		// check if file is an empty row
		if len(pieces) == 0 || (len(pieces) == 1 && pieces[0] == "") {
			return nil, errors.NewUtilsError("ParseMap", "invalid row in parsed file")
		}

		locationName := pieces[0]
		location := locations.GetByName(locationName)
		if location == nil {
			location = nodes.NewLocation(id, locationName)
			locations.Add(location)
		}
		for _, directionString := range pieces[1:] {
			var direction, neighboringLocationName string
			UnpackSliceString(strings.Split(directionString, "="), &direction, &neighboringLocationName)
			if err := validateDirection(direction, neighboringLocationName); err != nil {
				return nil, errors.NewUtilsErrorWrap("ParseDirections", err)
			}
			// check whether the location has already been added to the list of location
			neighboringLocation := locations.GetByName(neighboringLocationName)
			if neighboringLocation == nil {
				id++
				neighboringLocation = nodes.NewLocation(id, neighboringLocationName)
				locations.Add(neighboringLocation)
			}
			location.DirectionsOutBound.Add(direction, neighboringLocation, false)
			location.DirectionsInBound.Add(direction, neighboringLocation, false)
			neighboringLocation.DirectionsInBound.Add(direction, location, true)
		}
		id++
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.NewUtilsErrorWrap("ParseMap", err)
	}

	return locations, nil
}

func validateDirection(direction string, cityName string) error {
	if cityName == "" {
		return errors.NewUtilsError("parseDirections", "invalid city name as direction")
	}

	if direction == "" {
		return errors.NewUtilsError("parseDirections", "invalid direction")
	}
	if direction != constant.DirectionNorth &&
		direction != constant.DirectionWest &&
		direction != constant.DirectionSouth &&
		direction != constant.DirectionEast {
		return errors.NewUtilsError("parseDirections", fmt.Sprintf("uncataloged direction %s", direction))
	}
	return nil
}
