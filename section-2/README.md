## Section 2

### Useful Commands

Compiles your code to binary. Make sure that the folder structure is **src -> [packageName] -> ***
```
go install [packageName]
```

Enforces go formatting rules.
```
go fmt [packageName]
```

Golang has a great documentation that can be accessed on [doc](https://golang.org/). If you lack internet connection, golang can run locally an HTTP server that can be used.

```
godoc --http=localhost:6060/ &
```

```
go doc net/http Response | more
```