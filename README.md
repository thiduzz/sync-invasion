
# Alien Invasion (CodeKata)

Simulates an alien attack into cities defined by a text file - the loaded file should have the format descibed below:
```text
Foo  north=Bar  west=Baz  south=Qu-ux
Bar  south=Foo  west=Bee
```
The city name is the first string in the line and the remaining pieces are the name of the connecting cities prefixed with the direction name.

This are the set of rules are defined for the invasion and the world:
- Each city has a name and 4 directions (N,S,E,W) that might lead to different cities
- The simulation should enable the creation of N aliens
- Aliens are spawned in random cities
- Aliens will move randomly following the Directions of a city
- When an alien tries to move to a city which is already being attacked a fight happens
    - When fight happens both aliens die
    - When fight happens city is destroyed
    - Should print message "City X has been destroyed by Alien Y and Alien Z"
- When alien arrives in a city that has no connections it will be trapped
- Simulation ends when:
    - All aliens died
    - Or, each alien has moved for at least 10.000 times
- When simulation ends the map should be printed with the remaining cities in the same format as the input map file

## Pre-requisites

- Installed [Golang 1.18 or greater](https://go.dev/dl/)
- Installed [Mockgen 1.6.0](https://github.com/golang/mock) for Mock Generation

## Run Locally

Clone the project

```bash
  git clone https://github.com/thiduzz/invasion-sync
```

Go to the project directory

```bash
  cd invasion-sync
```

Build project

```bash
  go build ./cmd/cli/ -o ./bin/simulation
```

## Usage

```bash
     ./bin/simulation --map-path={MAP_FILE_PATH} --alien-qty={QUANTITY_OF_ALIENS} --max-moves={MAX_APPLICATION_ITERATION}
```
- `--map-path`: **(optional)** Path of file that represents the map of cities. Default to "./resources/world-map.txt"
- `--alien-qty`: **(optional)** Total amount of aliens to invade. Defaults to 100
- `--max-moves`: **(optional)** Amount of iterations all aliens will try to move. May be interrupted before, if simulation closing conditions are met. Defaults to 10000
-
### Example

```bash
     ./bin/simulation --map-path=./resources/world-map.txt --alien-qty=10 --max-moves=30
```
In the command above, a file located in the `resources` folder with the name `world-map.txt` will be loaded and define the cities. 10 aliens will be created. Each aliens will "move" at max 30 times between this cities.


## Authors

- [@thiduzz](https://www.github.com/thiduzz)


## License

[MIT](https://choosealicense.com/licenses/mit/)