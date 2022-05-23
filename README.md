# Tutorial

Go is a compiled language

compile and run:

    $ go run helloworld.go

Go natively handles Unicode

save the compiled result for later use:

    $ go build helloworld.go
    $ ./helloworld

Package `main` is special: It defines a standalone executable program, not a library.
Function `main` is also special: where the execution begins.

You must import *exactly* the packages you need.

You should use `gofmt` or `go fmt` and `goimports`

    $ go get golang.org/x/tools/cmd/goimports

---

    s := ""
    var s string
    var s = ""
    var s string = ""

