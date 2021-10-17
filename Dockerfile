FROM golang:1.17.2
WORKDIR /build
COPY go.* ./
RUN  go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /app ./cmd/pasted

FROM alpine:3.14.2
COPY --from=0 /app /usr/local/bin/pasted
VOLUME [ "/data" ]
LABEL org.opencontainers.image.source https://github.com/choopm/pasted
ENV HOST=[::]
ENV PORT=10101
ENV URL_ROOT=https://paste.0pointer.org/
CMD ["/usr/local/bin/pasted"]
