# Advanced Golang Programming - Go in Depth

## Build tool

Shows the coverage percentage of the package.

```
go test -cover [packegName]
```

Generates a report with the name `coverage.out` to be used by `go tool cover` command.

```
go test -coverprofile=coverage.out [packageName]
```

Analyze the report generated by `go test -coverprofile` command and shows the package level coverage.

```
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```