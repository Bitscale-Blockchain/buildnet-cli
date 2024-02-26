# Buildnet CLI

Buildnet CLI is a command-line interface (CLI) application for quickly scaffolding Cosmos blockchain projects based on provided configuration files.

## Installation

To install Buildnet CLI, you need to have Go installed on your system. Then, you can install it using `go get`:

```bash
go get -u github.com/yourusername/buildnet-cli
```

Replace `yourusername` with your GitHub username.

## Building

If you want to build the application from source, you can clone the repository and build it using Go:

```bash
git clone https://github.com/yourusername/buildnet-cli.git
cd buildnet-cli
go build
```

## Usage

Buildnet CLI provides several commands for creating, updating, and validating projects, as well as scaffolding new projects based on default configuration templates. Here are some examples of how to use the CLI:

- **Create a new project**: `buildnet-cli create config.json`
- **Update an existing project**: `buildnet-cli update config.json`
- **Validate a configuration file**: `buildnet-cli validate config.json`
- **Scaffold a new project**: `buildnet-cli scaffold defi --output myproject`
- **List available default configuration templates**: `buildnet-cli list-templates`
- **Use a specific default configuration template**: `buildnet-cli template defi --output myproject`

For more information on each command, you can use the `--help` flag, e.g. `buildnet-cli create --help`.

## Configuration File Format

The configuration file used by Buildnet CLI is in JSON format. Below is an example of the structure of the configuration file:

```json
{
  "name": "MyCosmosProject",
  "description": "A description of my Cosmos project",
  "modules": [
    {
      "name": "module1",
      "type": "moduleType1",
      "options": {
        "key1": "value1",
        "key2": "value2"
      }
    },
    {
      "name": "module2",
      "type": "moduleType2",
      "options": {
        "key1": "value1",
        "key2": "value2"
      }
    }
  ]
}
```

The configuration file consists of a top-level object with the following properties:

- `name`: The name of the project.
- `description`: A brief description of the project.
- `modules`: An array of modules that make up the project. Each module has a `name`, `type`, and `options`.
