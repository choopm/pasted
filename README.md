## Run the latest release using docker
```shell
docker run --rm -it -p 10101:10101 -v "$(pwd)/pastes:/data" ghcr.io/choopm/pasted
```

## Environment variables
* `HOST=[::]`
* `PORT=10101`
* `URL_ROOT=https://paste.0pointer.org/`

## Building local binaries
    # Current OS/Platform
    make

Binaries can be found in `dist/<platform>`

## Running a dev container
When using VS Code you may be asked to reopen the project in a devcontainer, do so.

Afterwards use `go run cmd/pasted/pasted.go` to run the app.
