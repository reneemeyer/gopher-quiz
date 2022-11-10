# Gopher Math Quiz

A terminal math game! Def not as cool as math blasters 

## Uses
- Go 1.19

## How to run

To use the default problems: 
`go build . && ./gopher-quiz -csv=problems.csv`

To use a custom problem set from another csv:

1. Import csv into root level of project. Csv must be in `question,answer` format  
2. `go build . && ./gopher-quiz -csv=<csv filename>`

## Tests

- coming soon, ok? relax!

