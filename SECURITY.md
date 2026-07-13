# Security Policy

## Supported Versions

We release patches for security vulnerabilities. The latest release is
always supported.

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |

## Reporting a Vulnerability

Please **do not** open a public GitHub issue for security vulnerabilities.

Instead, send an email to **my@dengmengmian.com** with:

- A clear description of the vulnerability
- Steps to reproduce
- Affected versions
- Any suggested fix (if available)

We will acknowledge your report within 72 hours and aim to publish a fix
within 30 days. We will credit you in the release notes unless you ask
us not to.

## Scope

This policy covers all packages under `github.com/dengmengmian/ghelper`,
including:

- `gconvert` — type conversion utilities
- `gstring` — string helpers (MD5, JSON, email validation)
- `gid` — unique ID generation (cryptographic and non-cryptographic)
- `gfile` — file system operations
- `ghttp` — HTTP client
- `gip` — client IP extraction
- `gmarkdown` — Markdown to HTML rendering
- `gos` — OS detection and home directory

## Disclosure Policy

Once a fix is released we will publish a GitHub Security Advisory. We
follow a coordinated disclosure process and will work with reporters to
agree on an appropriate timeline.
