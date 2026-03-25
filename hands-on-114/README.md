# hands-on-114

## Purpose
Practice from Go video lessons, followed by my own example implementation.

## What I Practiced
- Video concept:
- My own example:

## Run
```bash
go run .
```

## Notes
- `make([]string, 0, 50)` creates a slice with length `0` and capacity `50`, which helps avoid repeated reallocations while appending.
- `len(slice)` and `cap(slice)` are useful to verify slice growth behavior during practice.
- `append` can add many values in one call, which is useful for seeding sample data quickly.
- Iterating with `for i := 0; i < len(slice); i++` reinforces index-based access and ordered traversal.
- Writing my own version after the video helped me understand the difference between slice type, length, and capacity in a practical way.
