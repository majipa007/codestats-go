# codestats

A small Go CLI that scans a directory and prints line/character counts by file extension.

## Why
- Quick codebase size stats by language
- Terminal-first workflow
- Concurrency experiments (goroutines + semaphore)

## Build
```
go build -o codestats .
```

## Run
```
./codestats
./codestats .
./codestats /path/to/repo --workers=50
```

## Flags
- `--workers`: max concurrent goroutines (semaphore limit). Default is `500`.

## Config
The tool reads a JSON config file from:
- `$HOME/codestats/codestats.config.json`

Expected shape:
```
{
  "ignore_directories": [".git", "node_modules"],
  "allowed_extensions": [".go", ".py", ".js"]
}
```

## Perf runner
```
go run ./perf_runner
```

## Output
- Sorted by line count
- Totals and runtime included

## Notes
- Traversal is concurrent by directory and file.
- Concurrency is bounded by `--workers`.
