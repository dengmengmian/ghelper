# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Breaking
- `gid.GenShortID`: IDs now use the full 62-character alphabet (a-z, A-Z, 0-9)
  instead of being forcibly lowercased. This fixes the mismatch where the
  generator claimed a 62-char alphabet but `strings.ToLower` collapsed it to
  36 characters, halving the per-character entropy and making collisions far
  more likely than specified.

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
- `gid.Generate`: modulo bias in character selection eliminated via rejection
  sampling (previously `byte % 62` gave the first 8 characters ~25% higher
  probability than the rest).
- `gid`: collision tests rewritten to use a deterministic space large enough
  that the birthday-problem failure probability is < 1e-8 per CI run (was ~2.3%
  per test with the old 36-char effective alphabet).
- `gid.randChars`: fixed bug where it always returned an empty string.
- `gconvert.IsNumeric`: fixed false positive for strings with multiple dots (e.g., `"12.34.56"`).
- `Makefile`: corrected `VERSION_PACKAGE` to point to this repository.

## [0.1.0] - Initial Release

- Core packages: `gconvert`, `gstring`, `gid`, `gfile`, `ghttp`, `gip`, `gmarkdown`, `gos`.
