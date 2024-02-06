# godlp - Go Command-Line Interface for yt-dlp

<!-- ![godlp Logo](logo.png) -->

## Overview

godlp is a Go-based command-line interface (CLI) application that wraps commonly used functionality of the [yt-dlp](https://github.com/yt-dlp/yt-dlp) app. It leverages the [Cobra](https://github.com/spf13/cobra) framework for building CLI applications and [Viper](https://github.com/spf13/viper) for configuration management, providing a user-friendly and extensible interface for interacting with yt-dlp.

## Features

- **Wrapper for yt-dlp:** godlp simplifies the usage of yt-dlp by providing sensible defaults and helpful automations in it's CLI interface.
- **Configuration Management:** Utilizes Viper for easy configuration management, allowing users to customize settings through configuration files.
- **Extendable:** Easily extend and add new features by leveraging the modular nature of Cobra commands.

## Install

```bash
brew tap mcreekmore/mcreekmore
brew install godlp
```

## Develop

Make sure you have Go installed. Clone the godlp repository and build the executable:

```bash
git clone https://github.com/mcreekmore/godlp.git
cd godlp

# run
go run main.go

# build
go build -o godlp main.go

# run built binary
./godlp
```

## Usage

```
  godlp [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  soundcloud  For downloading tracks from soundcloud

Flags:
      --config string   config file (default is $HOME/.godlp.yaml)
  -h, --help            help for godlp
  -t, --toggle          help message for toggle

Use "godlp [command] --help" for more information about a command.
```

## Configuration

godlp uses Viper for configuration management. The configuration file is located at `~/.godlp.yaml` by default. You can customize settings in this file or create a new one using the `config` command.

Copy the `example.godlp.yaml` into your home directory.

```bash
cp example.godlp.yaml ~/.godlp.yaml
```

## License

This project is licensed under the Unlicense License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [yt-dlp](https://github.com/yt-dlp/yt-dlp): The underlying tool that godlp wraps.
- [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper): The frameworks used for building the CLI and managing configurations.
