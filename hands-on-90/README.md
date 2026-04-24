# hands-on-90

## Purpose
Practice the comma-ok idiom for safe map lookups.

## What I Practiced
- Reading values from a map.
- Checking whether a key exists with the `value, ok := map[key]` pattern.
- Printing different output depending on whether the key was present.

## Run
```bash
go run .
```

## Notes
- The program safely checks for keys in a map and reports whether each lookup succeeded.
- This folder is about map access safety, not slice iteration.
- The comma-ok pattern avoids confusing missing keys with zero values.
