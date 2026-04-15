# Understanding Writer Interfaces (Part 4)

## Purpose
Practice using the `io.Writer` interface with both files and in-memory buffers.

## What I Practiced
- Writing custom data to any type that satisfies `io.Writer`
- Sending the same data to an `os.File` and a `bytes.Buffer`
- Reusing one method for different output targets through interfaces

## Run
```bash
go run .
```

## Notes
- `Person.Writertest` writes the person's name to any supplied writer.
- Running the program creates or updates `new_file.txt` with `John Doe`.
- The same program also writes `hello` into a `bytes.Buffer` and prints the buffer contents.
