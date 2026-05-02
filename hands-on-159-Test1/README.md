Write a simple function and verify it with a unit test.

# Hands-on 159: First Unit Test

## Purpose
Practice writing a basic function and validating it with Go's testing package.

## What This Program Shows
- `add(a int, b int) int` returns the sum of two integers.
- `main()` prints the result of `add(2, 3)`.
- `main_test.go` checks that the function returns `5`.

## Run
```bash
go test .
go run .
```

## Expected Behavior
The test passes, and the program prints `5`.

## Key Takeaway
Small, focused tests are a good way to confirm simple behavior as you learn the language.
