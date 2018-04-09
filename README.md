# clcity
> CLI app

It's a CLI app for fetching and parsing data from 'Cuando Llega City Bus' API.

## Usage

1. Ensure you have go installed, if not, go [here](https://golang.org/doc/install)
2. Build and compile file `cmd/clcity/main.go`, you can do it both ways:
    1. Run `GOOS=linux go build -o ~/clcity cmd/clcity/main.go`, move `~/clcity` to a folder in the $PATH, use it as `clcity` everywhere.
    2. Run `go install`, it installs `clcity` in `$GOPATH/bin` folder, ensure you add it to the $PATH.
3. `clcity h` for detailed help.

## LICENSE

Under MIT, [read license](./LICENSE)
