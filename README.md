# Pokedex CLI

A command-line interface Pokedex application that interacts with the [PokeAPI](https://pokeapi.co/) to explore Pokemon locations and catch Pokemon.

## Features

- Interactive command-line interface
- Explore Pokemon location areas
- View Pokemon in specific locations
- Catch Pokemon with a chance-based system
- Built-in caching system to reduce API calls
- Pagination support for location areas

## Installation

### Prerequisites

- Go 1.16 or later

### Installing

1. Clone this repository:
```bash
git clone https://github.com/yourusername/pokedex-cli.git
```

2. Navigate to the project directory:
```bash
cd pokedex-cli
```

3. Build the application:
```bash
go build -o pokedex ./pokecli
```

## Usage

Run the application:
```bash
./pokedex
```

### Available Commands

- `help` - Display all available commands
- `exit` - Exit the application
- `map` - Show the next page of location areas
- `mapb` - Show the previous page of location areas
- `explore {location area}` - View Pokemon in a specific location area
- `catch {pokemon}` - Try to catch a specific Pokemon

## Examples

```
> help
Welcome to Pokedex CLI!
Available commands:
  help - Show this help message
  exit - Exit the REPL
  map - Show location areas
  mapb - Show previous location areas
  explore {location Area} - Explore a location area
  catch - Catch a Pokemon

> map
Location Areas:
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
...

> explore eterna-city-area
Exploring eterna-city-area
Pokemon in eterna-city-area
- budew
- roselia
- chansey
...

> catch pikachu
Caught!
Congratulations!
You have a new Pokemon!
Enjoy your new companion!
```

## Architecture

- Command-line interface built with Go's standard library
- Caching system to minimize API calls with configurable expiration
- HTTP client with timeout handling
- Modular command structure for easy extension

## License

[MIT License](LICENSE)
