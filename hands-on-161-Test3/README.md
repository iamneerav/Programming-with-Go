Build a formatted string and confirm it with a test.

# Hands-on 161: Testing Formatted Output

## Purpose
Practice testing a function that builds and returns a formatted string.

## What This Program Shows
- `location(x string) string` returns a sentence using `fmt.Sprintf()`.
- `main()` prints the formatted result for `Maldives`.
- `main_test.go` checks that the returned string matches the expected output.

## Run
```bash
go test .
go run .
```

## Expected Behavior
The test passes, and the program prints `I am in Maldives`.

## Key Takeaway
String-returning functions are easy to test because they have a clear input and output.
