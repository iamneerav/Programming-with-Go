Learn how a named return value can accumulate and return a sum.

# Hands-on 153: Named Return Values

## Purpose
Show how a named return variable can be used inside a function and returned at the end.

## What This Program Shows
- `namedreturn(x []int) (sum int)` declares `sum` in the function signature.
- The function loops through a slice of integers and updates `sum` directly.
- `main()` passes a slice and prints the total.
- A commented alternative shows the same logic without a named return.

## Run
```bash
go run .
```

## Expected Behavior
The program prints `15`.

## Key Takeaway
Named return values can make simple accumulator-style functions easier to read, but they should still be used carefully.
