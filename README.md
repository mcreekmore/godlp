# godlp - Go Command-Line Interface for yt-dlp

<!-- ![godlp Logo](logo.png) -->

## Overview

godlp is a Go-based command-line interface (CLI) application that wraps commonly used functionality of the [yt-dlp](https://github.com/yt-dlp/yt-dlp) app. It leverages the [Cobra](https://github.com/spf13/cobra) framework for building CLI applications and [Viper](https://github.com/spf13/viper) for configuration management, providing a user-friendly and extensible interface for interacting with yt-dlp.

## Features

- **Wrapper for yt-dlp:** godlp simplifies the usage of yt-dlp by providing sensible defaults and helpful automations in it's CLI interface.
- **Configuration Management:** Utilizes Viper for easy configuration management, allowing users to customize settings through configuration files.
- **Extendable:** Easily extend and add new features by leveraging the modular nature of Cobra commands.

## Installation

Make sure you have Go installed. Clone the godlp repository and build the executable:

```bash
git clone https://github.com/mcreekmore/godlp.git
cd godlp
go build -o godlp main.go
```

Run godlp:

```bash
./godlp
```

## Usage

## Configuration

godlp uses Viper for configuration management. The configuration file is located at `~/.godlp.yaml` by default. You can customize settings in this file or create a new one using the `config` command.

Example configuration file:

```yaml
# output_directory: ~/Downloads
# format: bestvideo+bestaudio
```

## License

This project is licensed under the Unlicense License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [yt-dlp](https://github.com/yt-dlp/yt-dlp): The underlying tool that godlp wraps.
- [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper): The frameworks used for building the CLI and managing configurations.
