# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added
- Comprehensive unit tests for all packages (`gconvert`, `gstring`, `gfile`, `ghttp`, `gip`, `gos`, `gmarkdown`, `gid`).
- GitHub Actions CI workflow for build, vet, and race-detector tests.
- `CONTRIBUTING.md` with contribution guidelines.
- Proper godoc comments on all exported functions and types.

### Changed
- `gstring.StringToJson` now returns `(map[string]interface{}, error)` instead of silently swallowing errors.
- `ghttp/tests`: all tests rewritten to use `httptest` (offline); removed `log.Panicln` calls.
- `gfile.IsDir`: removed noisy `fmt.Printf` debug output.
- `gfile.GetSize` / `ghttp.Body.FromToString`: updated from deprecated `ioutil.ReadAll` to `io.ReadAll`.

### Fixed
- `gid.randChars`: fixed bug where it always returned an empty string.
- `gconvert.IsNumeric`: fixed false positive for strings with multiple dots (e.g., `"12.34.56"`).
- `Makefile`: corrected `VERSION_PACKAGE` to point to this repository.

## [0.1.0] - Initial Release

- Core packages: `gconvert`, `gstring`, `gid`, `gfile`, `ghttp`, `gip`, `gmarkdown`, `gos`.
