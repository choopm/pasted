## Run using docker
```shell
docker run --rm -it -p 10101:10101 -v "$(pwd)/pastes:/data" registry.gitlab.0pointer.org/choopm/pasted
```

## Environment variables
* `HOST=[::]`
* `PORT=10101`
* `URL_ROOT=https://paste.0pointer.org/`