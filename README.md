# go-ls-ignorable

Simple `ls` command reimplementation with Golang, but filterable based on `.gitignore`

## How to use

### Build and run

```bash
go build -o build/
./build/go-ls-ignorable.exe -g ./example/gitignore -d ./resources
```

### Parameters

```bash
./build/go-ls-ignorable.exe --help
```

```text                   
NAME:
   go-ls-ignorable - Simple ls command reimplementation, but filterable with .gitignore

USAGE:
   go-ls-ignorable [global options]

GLOBAL OPTIONS:
   --gitignore value, -g value  Path to .gitignore file
   --dir value, -d value        Path to directory to list files (default: ".")
   --help, -h                   show help
```

* `--gitignore (-g)`: Path to .gitignore
* `--dir (-d)`: Path to directory to list files

## License

Apache-2.0
