Show the execution order of deferred function calls.

# Hands-on 156: Defer Order

## Purpose
Demonstrate that deferred calls run after the surrounding function finishes and in last-in, first-out order.

## What This Program Shows
- `main()` defers five `fmt.Println()` calls.
- The normal `fmt.Println("Hello")` runs before any deferred call.
- The deferred values are printed in reverse order: `5` back down to `1`.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `Hello` first, then `5`, `4`, `3`, `2`, and `1`.

## Key Takeaway
Deferred calls are stacked and executed in reverse order when the function returns.
