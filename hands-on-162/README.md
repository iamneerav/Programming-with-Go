Return a descriptive sentence from a simple function.

# Hands-on 162: Simple Return Function

## Purpose
Show how a function can accept one value and return a human-readable string.

## What This Program Shows
- `Paradise(loc string) string` builds a sentence using `fmt.Sprint()`.
- `main()` calls the function with `Hawaii` and prints the result.
- The function name is exported to illustrate normal Go naming rules.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `My idea of paradise is Hawaii`.

## Key Takeaway
Even a simple helper function is a good place to practice parameters, return values, and string formatting.
