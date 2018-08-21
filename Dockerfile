FROM gitlab.exservice.io:4567/container-registry/golang:1.8

COPY ./ /go/src/cabal-api

WORKDIR /go/src/cabal-api

# Get dependency
RUN go get -v && \
    go build && \
    chmod -R 777 ./cabal-api

EXPOSE 90

ENTRYPOINT ["./cabal-api"]