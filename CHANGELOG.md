# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive test suite using testify framework
- Test coverage for tilde expansion, path resolution, and directory creation
- Refactored code into testable functions for better maintainability

### Changed
- Build process now uses goreleaser for local builds instead of direct `go build`
- Improved code structure with extracted functions: `expandTilde()`, `resolvePath()`, `validateOrCreateDir()`, and `runMkcd()`

## [0.1.0] - YYYY-MM-DD

### Added
- Initial release of mkcd tool
- Create directories with nested parent directories (mkdir -p behavior)
- Change directory functionality
- Support for tilde (~) expansion for home directory
- Cross-platform support (Windows, macOS, Linux)
- Support for both relative and absolute paths
- Error handling for existing files (when path exists but is not a directory)
- Validation of directory existence before creation

[Unreleased]: https://github.com/Naven/mkcd/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/Naven/mkcd/releases/tag/v0.1.0
