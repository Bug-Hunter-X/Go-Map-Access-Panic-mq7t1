# Go Map Access Panic

This repository demonstrates a common error in Go: panics caused by accessing a map without checking for `nil`.  Accessing a `nil` map will result in a runtime panic. 

The `bug.go` file shows the problematic code.  The `bugSolution.go` file demonstrates a safe solution.

## How to Run

1. Clone this repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go` (This will panic).
4. Run `go run bugSolution.go` (This will handle the nil case gracefully).
