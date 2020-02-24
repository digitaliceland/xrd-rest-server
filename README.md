# rest-server

A REST test webservice to test X-Road services

## Getting started
You can download compiled binaries for your architecture:

* [Linux x64](release/linux/rest-server)
* [MacLOS x64](release/macos/rest-server)
* [Windows x64](release/windows/rest-server.exe)

## Usage examples
Then run the script to bind to the port requested.  E.g. for localhost, port 1234 please use:

```console
$ ./rest-server --addr 127.0.0.1:1234
```


To bind to all IP addresses on the box on port 7777:
```console
$ ./rest-server --addr :7777
```

By default it binds to :1234


## Debugging

You can run code directly with go like this:
```console
$ go run main.go
```

## Building

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`

Then to compile and run it just:

```console
$ go build main.go
$ ./main
```

To build for multiple architectures we also provide a build script, build.sh
