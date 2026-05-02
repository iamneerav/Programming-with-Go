Test function values passed into another function.

# Hands-on 160: Functions as Values

## Purpose
Show that functions can be passed as arguments and tested like any other behavior.

## What This Program Shows
- `add` and `sub` share the same function signature.
- `doMath(a, b, f)` accepts a function and calls it.
- `main()` prints the types and runs both operations.
- `main_test.go` verifies `add` and `sub` with unit tests.

## Run
```bash
go test .
go run .
```

## Expected Behavior
The tests pass, and the program prints the function types followed by `10` and `6`.

## Key Takeaway
Go treats functions as first-class values, which makes callbacks and reusable behavior straightforward.
