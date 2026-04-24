# Hands-on 152: Wrapper Functions

## Purpose
Show how a wrapper function can add behavior around existing functionality without changing the original implementation.

## What This Program Shows
- The active example uses `writeLog(s fmt.Stringer)` as a wrapper around logging.
- `book` and `count` both implement `String() string`, so both satisfy `fmt.Stringer`.
- `writeLog()` accepts either type and logs the string representation.
- The commented second example shows error wrapping with `readFile()` calling `os.ReadFile()` and returning a clearer error message with `fmt.Errorf()`.

## Run
```bash
go run .
```

## Expected Behavior
The active program logs the formatted `book` and `count` values. The commented example is included as study material for wrapping file-read errors.

## Key Takeaway
Wrapper functions are useful for reusable cross-cutting behavior such as logging and error handling.
