Use a closure to preserve state across repeated function calls.

# Hands-on 168: Closures

## Purpose
Demonstrate how a closure can keep track of changing state between calls.

## What This Program Shows
- `powinator(a)` returns a function that remembers the exponent counter `c`.
- Each call increments `c` and returns the next power of `a`.
- `main()` calls the returned function multiple times to show the sequence.

## Run
```bash
go run .
```

## Expected Behavior
The program prints successive powers of `2`: `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`.

## Key Takeaway
Closures retain access to variables from their surrounding scope even after the outer function has returned.
