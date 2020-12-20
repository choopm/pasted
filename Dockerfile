FROM golang
ARG CI_PROJECT_DIR=/builds/pasted
ENV CI_PROJECT_DIR=$CI_PROJECT_DIR
ARG CI_PROJECT_URL=https://gitlab.0pointer.org/choopm/pasted
ENV CI_PROJECT_URL=$CI_PROJECT_URL
RUN export REPO_NAME=`echo $CI_PROJECT_URL|sed 's/.*:\/\///g;'` && \
    mkdir -p $GOPATH/src/$(dirname $REPO_NAME) && \
    ln -svf $CI_PROJECT_DIR $GOPATH/src/$REPO_NAME
WORKDIR $CI_PROJECT_DIR
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app ./cmd/pasted

FROM alpine
WORKDIR /opt/pasted
COPY --from=0 /app .
VOLUME [ "/data" ]
ENV HOST=[::]
ENV PORT=10101
ENV URL_ROOT=https://paste.0pointer.org/
CMD ["./app"]
