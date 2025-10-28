# mkcd

A cross-platform CLI tool that combines `mkdir` and `cd` functionality, allowing you to create a directory and change into it in a single command.

## Features

- Creates directories (including parent directories if needed)
- Validates directory existence and handles errors gracefully
- Cross-platform support (Windows, macOS, Linux)
- Handles both relative and absolute paths
- Supports tilde (`~`) expansion for home directory

## Installation

```sh
go install github.com/Naven/mkcd@latest
```

## Usage

### Basic Shell Integration

Since a child process cannot change the parent shell's directory, you need to use one of these methods:

#### Bash/Zsh (Recommended)
Add this function to your `~/.bashrc` or `~/.zshrc`:

```bash
mkcd() {
    local dir
    dir=$(command mkcd "$@")
    [ -n "$dir" ] && cd "$dir"
}
```

Then use it like:
```sh
mkcd foo           # Creates `foo` and changes into it
mkcd ~/projects/new # Creates directory in home/projects and changes into it
mkcd path/to/deep  # Creates nested directories and changes into the last one
```

#### POSIX Shell
```sh
mkcd() {
    dir=$(command mkcd "$@")
    [ -n "$dir" ] && cd "$dir"
}
```

#### PowerShell (Windows)
```powershell
function mkcd {
    $dir = mkcd.exe $args
    if ($dir) { Set-Location $dir }
}
```

### Direct Usage

You can also use it directly and manually cd:
```sh
cd $(mkcd foo)
```
