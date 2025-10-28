# mkcd

A cross-platform CLI tool that combines `mkdir` and `cd` functionality, allowing you to create a directory and change into it in a single command.

## Features

- Creates directories (including parent directories if needed)
- Validates directory existence and handles errors gracefully
- Cross-platform support (Windows, macOS, Linux)
- Handles both relative and absolute paths
- Supports tilde (`~`) expansion for home directory

## Usage

Use it like:
```sh
mkcd foo           # Creates `foo` and changes into it
mkcd ~/projects/new # Creates directory in home/projects and changes into it
mkcd path/to/deep  # Creates nested directories and changes into the last one
```
