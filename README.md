# 06-November-2026

This readme contains the required set of terminologies and concepts that one should be familiar with to facilitate the *entire* learning process. See [here](comments.md) for why.


## Anatomy of a Go project
- `Source code/file`: a file with `.go` extension and a `package <>` directive at the top, e.g., `package controllers`, `package repositories`, etc.
- `Package`: a folder/directory containing 1 or more source files.
- `Module`: a folder containing one or more packages defined by a `go.mod` file. 




## go.mod
- As explained, one **must always** create at least 1 go module in order to write go code. See [here](go_module.md) for why.
- It can be created by hand or with the command: `go mod init module_name`
- Modules and packages should *preferably* be in directories bearing their names.


## go.work

- It's sometimes necessary to use a `go.work` file.
- It can be created by hand or with the command: `go work init` 
- See [example](go.work): `use` lets the `go` command find everything it needs.

