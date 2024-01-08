# qalogparser - A Quake Arena Logger parser

![Logo](assets/logo_quakearena.jpg)

`qaparser` is a command-line tool for parsing Quake Arena log files. It's designed to extract and analyze gameplay data efficiently, offering both JSON and plain text output formats.

## Usage

### Getting Help

To get a list of available commands and global flags:

```bash
qaparser --help

A Quake Arena log parser

Usage:
  qaparser [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  parser      A parser for Quake Arena log files

Flags:
  -h, --help     help for qaparser
  -t, --toggle   Help message for toggle

Use "qaparser [command] --help" for more information about a command.
```

### Parsing Quake Arena Log Files

To parse log files:

```bash
qaparser parser --help

A parser for Quake Arena log files

Usage:
  qaparser parser [flags]

Flags:
  -h, --help           help for parser
  -i, --input string   Input file
  -k, --kind string    Kind of output (json, text) (default "json")
```

Use case

```bash
qaparser parser -i=/quake/log/quakeqa.log
```

The app going to print to stdout, so if you need to save the output to a file, just drop the content to output, just like that.

```bash
qaparser parser -i=/quake/log/quakeqa.log > myfile.json
```

or a text output (formated)

```bash
qaparser parser -i=/quake/log/quakeqa.log  -k=text > myfile.txt
```

### Build the project

Just run the command `make build`

```bash
make build
```

The binary file going to folder `build`, you can move to a expoted path for better use.

or `make help` for more details

```bash
make help
```
-----

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-qalogparser/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE.md) file.

**2024**, thiagozs
