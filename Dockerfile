FROM --platform=${BUILDPLATFORM} golang:1.19.3 as base

FROM base as builder
ARG GITHUB_TOKEN=
ENV GOPRIVATE=github.com/choopm/*
RUN git config --global \
        url."https://${GITHUB_TOKEN}@github.com/choopm".insteadOf \
        "https://github.com/choopm"
WORKDIR /build
COPY go.* ./
RUN  go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /pasted ./cmd/pasted

# For building local binaries via make
FROM base AS baselocalbuilder
RUN apt update && apt -y upgrade && apt install -y openssh-client

FROM baselocalbuilder AS localbuilder
RUN mkdir /build
RUN mkdir -p -m 0600 ~/.ssh && \
        ssh-keyscan github.com >> ~/.ssh/known_hosts
RUN git config --global \
        url."git@github.com:choopm/".insteadOf \
        "https://github.com/choopm/"
ENV GOPRIVATE=github.com/choopm/*
WORKDIR /build
COPY go.* .
RUN --mount=type=ssh go mod download
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=ssh cd /build/cmd/pasted && \
        GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build
FROM scratch as bin-unix
COPY --from=localbuilder /build/cmd/pasted/pasted /
FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin
FROM scratch AS bin-windows
COPY --from=localbuilder /build/cmd/pasted/pasted.exe /
FROM bin-${TARGETOS} AS bin

# default target
FROM alpine:3.14.2
COPY --from=builder /pasted /usr/local/bin/pasted
VOLUME [ "/data" ]
LABEL org.opencontainers.image.source https://github.com/choopm/pasted
ENV HOST=[::]
ENV PORT=10101
ENV URL_ROOT=https://paste.0pointer.org/
CMD ["/usr/local/bin/pasted"]
