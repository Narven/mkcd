# DEVELOPMENT


## Development

This project uses [Taskfile](https://taskfile.dev) for build automation. Install Task with:

```sh
# macOS/Linux
brew install go-task/tap/go-task

# Or using the install script
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin
```

### Available Tasks

```sh
task                    # List all available tasks
task build              # Build the binary
task build-release      # Build optimized binary for release
task run                # Build and run the tool
task install            # Install to $GOPATH/bin or $GOBIN
task test               # Run tests
task test-coverage      # Run tests with coverage report
task fmt                # Format Go code
task vet                # Run go vet
task lint               # Run fmt and vet
task check              # Run all checks (fmt, vet, test)
task clean              # Clean build artifacts
task cross-build        # Cross-compile for multiple platforms
task version            # Show Go version and module info
```
