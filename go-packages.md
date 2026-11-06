- Rule number 1 with packages: every Go code that you write must be in a Go package. That's just a corrolary from the fact that every go source code such as <file_name>.go is in 1 Go package. 


If you try to execute a go file with a package directive other than `main`, you would get the following error:

```
└─❯ go run test.go 
package command-line-arguments is not a main package
```

You can only run: `go run <file_name>.go` successfuly iff `<file_name>.go` contains the `package main` directive.