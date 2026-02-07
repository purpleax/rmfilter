# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`rmfilter` is a Windows 11 CLI utility written in Go that recursively searches a directory tree and deletes files whose names contain a given substring pattern. It confirms before deleting unless `--force` is passed.

## Build and Run

```bash
go build -o rmfilter.exe .
```

Usage: `rmfilter [--force] <directory> <pattern>`

There are no third-party dependencies (stdlib only), no tests, and no linter configured.

## Architecture

Single-file program (`main.go`) — all logic is in `func main()`:
1. Parses flags (`--force`) and positional args (directory, pattern)
2. Walks the directory with `filepath.Walk`, collecting files whose `Name()` contains the pattern
3. Lists matches and prompts for confirmation (skipped with `--force`)
4. Deletes confirmed files and reports count
