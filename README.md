# rmfilter

A command-line utility that searches a directory for files whose names contain a given pattern and deletes them.

## Build

```
go build -o rmfilter.exe .
```

## Usage

```
rmfilter [options] <directory> <pattern>
```

### Examples

```bash
# Delete files containing "_hls_" in C:\files (top-level only, with confirmation)
rmfilter C:\files "_hls_"

# Include subdirectories
rmfilter --recursive C:\files "_hls_"

# Skip confirmation
rmfilter --force C:\files "_hls_"

# Preview without deleting
rmfilter --dry-run --recursive C:\files ".tmp"

# See each file as it's deleted
rmfilter --verbose --force C:\files "_hls_"
```

### Flags

| Flag | Description |
|------|-------------|
| `--force` | Skip confirmation prompt and delete immediately |
| `--dry-run` | List matching files without deleting |
| `--recursive` | Search subdirectories recursively |
| `--verbose` | Show each file as it is deleted |
