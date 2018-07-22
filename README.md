# clcity

> CLI app

It's a CLI app for fetching and parsing data from 'Cuando Llega City Bus' API.

## Requirements

- Go: `> 1.9`
- Dep: `> 0.4`

## Usage

1.  Install this repo:

        go get -u github.com/cafrias/clcity

2.  Go to repo too and install dependencies:

        cd ~/go/src/github.com/cafrias/clcity
        dep ensure

3.  Build and compile file `cmd/clcity/main.go`, you can do it both ways:
    1.  Run `GOOS=linux go build -o ~/clcity cmd/clcity/main.go`, move `~/clcity` to a folder in the $PATH, use it as `clcity` everywhere.
    2.  Run `go install`, it installs `clcity` in `$GOPATH/bin` folder, ensure you add it to the $PATH.
4.  `clcity h` for detailed help.

## LICENSE

Under MIT, [read license](./LICENSE)
